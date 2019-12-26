package allure

import (
	"encoding/json"
	"fmt"
	"github.com/dailymotion/allure-go/parameter"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
)

//result is the top level report object for a test
type Result struct {
	UUID          string                `json:"uuid,omitempty"`
	Name          string                `json:"name,omitempty"`
	Description   string                `json:"description,omitempty"`
	Status        string                `json:"status,omitempty"`
	StatusDetails *StatusDetails        `json:"statusDetails,omitempty"`
	Stage         string                `json:"stage,omitempty"`
	Steps         []StepObject          `json:"steps,omitempty"`
	Attachments   []Attachment          `json:"attachments,omitempty"`
	Parameters    []parameter.Parameter `json:"parameters,omitempty"`
	Start         int64                 `json:"start,omitempty"`
	Stop          int64                 `json:"stop,omitempty"`
	Children      []string              `json:"children,omitempty"`
	Befores       []Before              `json:"befores,omitempty"`
	FullName      string                `json:"fullName,omitempty"`
	Labels        []Label               `json:"labels,omitempty"`
	Test          func()
}
type FailureMode string

//Before defines a step
type Before struct {
	Name          string         `json:"name,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetails *StatusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage,omitempty"`
	Description   string         `json:"description,omitempty"`
	Start         int64          `json:"start,omitempty"`
	Stop          int64          `json:"stop,omitempty"`
	Steps         []StepObject   `json:"steps,omitempty"`
	Attachments   []Attachment   `json:"attachments,omitempty"`
}

// This interface provides functions required to manipulate children step records, used in the result object and
// step object for recursive handling
type hasSteps interface {
	GetSteps() []StepObject
	AddStep(step StepObject)
}

type hasAttachments interface {
	GetAttachments() []Attachment
	AddAttachment(attachment Attachment)
}

type hasStatus interface {
	SetStatus(status string)
	GetStatus() string
}

func (r *Result) GetAttachments() []Attachment {
	return r.Attachments
}

func (r *Result) AddAttachment(attachment Attachment) {
	r.Attachments = append(r.Attachments, attachment)
}

func (r *Result) GetSteps() []StepObject {
	return r.Steps
}

func (r *Result) AddStep(step StepObject) {
	r.Steps = append(r.Steps, step)
}

func (r *Result) SetStatus(status string) {
	r.Status = status
}

func (r *Result) GetStatus() string {
	return r.Status
}

func (r *Result) setLabels(t *testing.T, labels TestLabels) {
	wsd := os.Getenv(WsPathEnvKey)

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

	r.AddLabel("package", testPackage)
	r.AddLabel("testClass", testPackage)
	r.AddLabel("testMethod", t.Name())
	if hostname, err := os.Hostname(); err == nil {
		r.AddLabel("host", hostname)
	}

	r.AddLabel("language", "golang")

	//TODO: these labels are available, but should be handled separately.

	//	ParentSuite string
	//	Suite       string
	//	SubSuite    string
	//	Thread      string
	//	Framework   string
}

func (r *Result) setDefaultLabels(t *testing.T) {
	wsd := os.Getenv(WsPathEnvKey)

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

	r.AddLabel("package", testPackage)
	r.AddLabel("testClass", testPackage)
	r.AddLabel("testMethod", t.Name())
	if hostname, err := os.Hostname(); err == nil {
		r.AddLabel("host", hostname)
	}

	r.AddLabel("language", "golang")

	//TODO: these labels are available, but should be handled separately.

	//	ParentSuite string
	//	Suite       string
	//	SubSuite    string
	//	Thread      string
	//	Framework   string
}

func (r *Result) AddLabel(name string, value string) {
	r.Labels = append(r.Labels, Label{
		Name:  name,
		Value: value,
	})
}

func (r *Result) writeResultsFile() error {
	CreateFolderOnce.Do(createFolderIfNotExists)
	CopyEnvFileOnce.Do(copyEnvFileIfExists)

	j, err := json.Marshal(r)
	if err != nil {
		return errors.Wrap(err, "Failed to marshall result into JSON")
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s-result.json", ResultsPath, r.UUID), j, 0777)
	if err != nil {
		return errors.Wrap(err, "Failed to write in file")
	}
	return nil
}

func newResult() *Result {
	return &Result{
		UUID:  GenerateUUID(),
		Start: getTimestampMs(),
	}
}

func getTimestampMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func createFolderIfNotExists() {
	resultsPathEnv := os.Getenv(ResultsPathEnvKey)
	if resultsPathEnv == "" {
		log.Printf("environment variable %s cannot be empty\n", ResultsPathEnvKey)
	}
	ResultsPath = fmt.Sprintf("%s/allure-results", resultsPathEnv)

	if _, err := os.Stat(ResultsPath); os.IsNotExist(err) {
		err = os.Mkdir(ResultsPath, 0777)
		if err != nil {
			log.Println(err, "Failed to create allure-results folder")
		}
	}
}

func copyEnvFileIfExists() {
	if envFilePath := os.Getenv(EnvFileKey); envFilePath != "" {
		envFilesStrings := strings.Split(envFilePath, "/")
		if ResultsPath != "" {
			if _, err := copy(envFilePath, ResultsPath+"/"+envFilesStrings[len(envFilesStrings)-1]); err != nil {
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
			log.Printf("Could not close the stream for the environment file, %f\n", err)
		}
	}()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err = destination.Close(); err != nil {
			log.Printf("Could not close the stream for the destination of the environment file, %f\n", err)
		}
	}()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
