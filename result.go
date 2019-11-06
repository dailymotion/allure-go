package allure

import (
	"encoding/json"
	"fmt"
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

type hasAttachments interface {
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

func (r *result) setLabels(t *testing.T) {
	wsd := os.Getenv(wsPathEnvKey)

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
	resultsPathEnv := os.Getenv(resultsPathEnvKey)
	if resultsPathEnv == "" {
		log.Fatalf("%s environment variable cannot be empty", resultsPathEnvKey)
	}
	resultPath = fmt.Sprintf("%s/allure-results", resultsPathEnv)

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
		Start: getTimestampMs(),
	}
}

func getTimestampMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
