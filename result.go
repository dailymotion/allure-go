package allure

import (
	"encoding/json"
	"fmt"
	"github.com/jtolds/gls"
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
type result struct {
	UUID          string         `json:"uuid,omitempty"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetails *statusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage,omitempty"`
	Steps         []stepObject   `json:"steps,omitempty"`
	Attachments   []Attachment   `json:"attachments,omitempty"`
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
	Attachments   []Attachment   `json:"attachments,omitempty"`
}

type HasSteps interface {
	GetSteps() []stepObject
	AddStep(step stepObject)
}

func (r *result) GetSteps() []stepObject {
	return r.Steps
}

func (r *result) AddStep(step stepObject) {
	r.Steps = append(r.Steps, step)
}

type Parameter struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

var wsd, resultPath string

const (
	ALLURE_RESULTS_PATH = "ALLURE_RESULTS_PATH"
	nodeKey             = "current_step_container"
)

//Test execute the test and creates an Allure result used by Allure reports
func Test(t *testing.T, description string, testFunc func()) {
	wsd = os.Getenv(ALLURE_RESULTS_PATH)
	if wsd == "" {
		log.Fatalf(fmt.Sprintf("%s environment variable cannot be empty", ALLURE_RESULTS_PATH))
		os.Exit(1)
	}
	resultPath = fmt.Sprintf("%s/allure-results", wsd)

	var r *result
	r = newResult()
	r.UUID = generateUUID()
	r.Start = getTimestampMs()
	r.Name = t.Name()
	r.Description = description
	r.setLabels(t)
	r.Steps = make([]stepObject, 0)

	defer func() {
		r.Stop = getTimestampMs()
		r.Status = getTestStatus(t)
		r.Stage = "finished"

		err := r.writeResultsFile()
		//err := r.writeResultsFile()
		if err != nil {
			log.Fatalf(fmt.Sprintf("Failed to write content of result to json file"), err)
			os.Exit(1)
		}
	}()
	ctxMgr.SetValues(gls.Values{"test_result_object": r, nodeKey: r}, testFunc)
}

func (r *result) setLabels(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(2)
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
	j, err := json.Marshal(r)
	if err != nil {
		return errors.Wrap(err, "Failed to marshall result into JSON")
	}
	if _, err := os.Stat(resultPath); os.IsNotExist(err) {
		err = os.Mkdir(resultPath, 0777)
		if err != nil {
			return errors.Wrap(err, "Failed to create allure-results folder")
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s-result.json", resultPath, r.UUID), j, 0777)
	if err != nil {
		return errors.Wrap(err, "Failed to write in file")
	}
	return nil
}

func newResult() *result {
	return &result{
		UUID:  generateUUID(),
		Start: time.Now().Unix(),
	}

func getTimestampMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
