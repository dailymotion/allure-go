package allure

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/jtolds/gls"
	"github.com/pkg/errors"
)

//result is the top level report object for a test
type result struct {
	UUID          string         `json:"uuid,omitempty"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetails *statusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage,omitempty"`
	Steps         []stepObject   `json:"steps,omitempty"`
	Attachments   []attachment   `json:"attachments,omitempty"`
	Parameters    []Parameter    `json:"parameters,omitempty"`
	Start         int64          `json:"start,omitempty"`
	Stop          int64          `json:"stop,omitempty"`
	Children      []string       `json:"children,omitempty"`
	Befores       []Before       `json:"befores,omitempty"`
	FullName      string         `json:"fullName,omitempty"`
	Labels        []Label        `json:"labels,omitempty"`
}
type FailureMode string

//Before defines a step
type Before struct {
	Name          string         `json:"name,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetails *statusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage,omitempty"`
	Description   string         `json:"description,omitempty"`
	Start         int64          `json:"start,omitempty"`
	Stop          int64          `json:"stop,omitempty"`
	Steps         []stepObject   `json:"steps,omitempty"`
	Attachments   []attachment   `json:"attachments,omitempty"`
}

// This interface provides functions required to manipulate children step records, used in the result object and
// step object for recursive handling
type hasSteps interface {
	GetSteps() []stepObject
	AddStep(step stepObject)
}

type HasAttachments interface {
	GetAttachments() []attachment
	AddAttachment(attachment attachment)
}

func (r *result) GetAttachments() []attachment {
	return r.Attachments
}

func (r *result) AddAttachment(attachment attachment) {
	r.Attachments = append(r.Attachments, attachment)
}

func (r *result) GetSteps() []stepObject {
	return r.Steps
}

func (r *result) AddStep(step stepObject) {
	r.Steps = append(r.Steps, step)
}

// TestWithParameters executes a test and adds parameters to the Allure result object
func TestWithParameters(t *testing.T, description string, parameters map[string]interface{}, testFunc func()) {
	var r *result
	r = newResult()
	r.UUID = generateUUID()
	r.Start = getTimestampMs()
	r.Name = t.Name()
	r.Description = description
	r.setLabels(t)
	r.Steps = make([]stepObject, 0)
	if parameters == nil || len(parameters) > 0 {
		r.Parameters = convertMapToParameters(parameters)
	}

	defer func() {
		r.Stop = getTimestampMs()
		r.Status = getTestStatus(t)
		r.Stage = "finished"

		err := r.writeResultsFile()
		if err != nil {
			log.Fatalf(fmt.Sprintf("Failed to write content of result to json file"), err)
			os.Exit(1)
		}
	}()
	ctxMgr.SetValues(gls.Values{"test_result_object": r, nodeKey: r}, testFunc)
}

//Test execute the test and creates an Allure result used by Allure reports
func Test(t *testing.T, description string, testFunc func()) {
	TestWithParameters(t, description, nil, testFunc)
}

func (r *result) setLabels(t *testing.T) {
	wsd := os.Getenv(wsPathEnvKey)

	programCounters := make([]uintptr, 10)
	callersCount := runtime.Callers(0, programCounters)
	var testFile string
	for i := 0; i < callersCount; i++ {
		_, testFile, _, _ = runtime.Caller(i)
		if strings.Contains(testFile, "_test.go") {
			break
		}
	}
	testPackage := strings.TrimSuffix(strings.Replace(strings.TrimPrefix(testFile, wsd+"/"), "/", ".", -1), ".go")

	pkgLabel := Label{
		Name:  "package",
		Value: testPackage,
	}
	r.Labels = append(r.Labels, pkgLabel)
	classLabel := Label{
		Name:  "testClass",
		Value: testPackage,
	}
	r.Labels = append(r.Labels, classLabel)
	methodLabel := Label{
		Name:  "testMethod",
		Value: t.Name(),
	}
	r.Labels = append(r.Labels, methodLabel)
}

func (r *result) writeResultsFile() error {
	createFolderOnce.Do(createFolderIfNotExists)
	copyEnvFileOnce.Do(copyEnvFileIfExists)

	j, err := json.Marshal(r)
	if err != nil {
		return errors.Wrap(err, "Failed to marshall result into JSON")
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s-result.json", resultsPath, r.UUID), j, 0777)
	if err != nil {
		return errors.Wrap(err, "Failed to write in file")
	}
	return nil
}

func newResult() *result {
	return &result{
		UUID:  generateUUID(),
		Start: getTimestampMs(),
	}
}

func getTimestampMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func createFolderIfNotExists() {
	resultsPathEnv := os.Getenv(resultsPathEnvKey)
	if resultsPathEnv == "" {
		log.Printf("environment variable %s cannot be empty\n", resultsPathEnvKey)
	}
	resultsPath = fmt.Sprintf("%s/allure-results", resultsPathEnv)

	if _, err := os.Stat(resultsPath); os.IsNotExist(err) {
		err = os.Mkdir(resultsPath, 0777)
		if err != nil {
			log.Println(err, "Failed to create allure-results folder")
		}
	}
}

func copyEnvFileIfExists() {
	if envFilePath := os.Getenv(envFileKey); envFilePath != "" {
		envFilesStrings := strings.Split(envFilePath, "/")
		if resultsPath != "" {
			if _, err := copy(envFilePath, resultsPath+"/"+envFilesStrings[len(envFilesStrings)-1]); err != nil {
				log.Println("Could not copy the environment file", err)
			}
		}

	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err = source.Close(); err != nil {
			log.Fatalf("Could not close the stream for the environment file, %f", err)
		}
	}()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err = destination.Close(); err != nil {
			log.Fatalf("Could not close the stream for the destination of the environment file, %f", err)
		}
	}()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
