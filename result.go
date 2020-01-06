package allure

import (
	"encoding/json"
	"fmt"
	"github.com/dailymotion/allure-go/parameter"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"
)

//result is the top level report object for a test
type Result struct {
	UUID          string                `json:"uuid,omitempty"`
	Name          string                `json:"name,omitempty"`
	Description   string                `json:"description,omitempty"`
	Status        string                `json:"status,omitempty"`
	StatusDetails *statusDetails        `json:"statusDetails,omitempty"`
	Stage         string                `json:"stage,omitempty"`
	Steps         []StepObject          `json:"steps,omitempty"`
	Attachments   []Attachment          `json:"attachments,omitempty"`
	Parameters    []parameter.Parameter `json:"parameters,omitempty"`
	Start         int64                 `json:"start,omitempty"`
	Stop          int64                 `json:"stop,omitempty"`
	Children      []string              `json:"children,omitempty"`
	FullName      string                `json:"fullName,omitempty"`
	Labels        []Label               `json:"labels,omitempty"`
	Test          func()                `json:"-"`
}

func (r *Result) addReason(reason string) {
	testStatusDetails := r.StatusDetails
	if testStatusDetails == nil {
		testStatusDetails = &statusDetails{}
	}
	r.StatusDetails.Message = reason
}

func (r *Result) addDescription(description string) {
	r.Description = description
}

func (r *Result) addParameter(name string, value interface{}) {
	r.Parameters = append(r.Parameters, parseParameter(name, value))
}

func (r *Result) addName(name string) {
	r.Name = name
}

func (r *Result) addAction(action func()) {
	r.Test = action
}

type FailureMode string

func (r *Result) getAttachments() []Attachment {
	return r.Attachments
}

func (r *Result) addAttachment(attachment Attachment) {
	r.Attachments = append(r.Attachments, attachment)
}

func (r *Result) getSteps() []StepObject {
	return r.Steps
}

func (r *Result) addStep(step StepObject) {
	r.Steps = append(r.Steps, step)
}

func (r *Result) setStatus(status string) {
	r.Status = status
}

func (r *Result) getStatus() string {
	return r.Status
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

	r.addLabel("package", testPackage)
	r.addLabel("testClass", testPackage)
	r.addLabel("testMethod", t.Name())
	if hostname, err := os.Hostname(); err == nil {
		r.addLabel("host", hostname)
	}

	r.addLabel("language", "golang")

	//TODO: these labels are available, but should be handled separately.

	//	ParentSuite string
	//	Suite       string
	//	SubSuite    string
	//	Thread      string
	//	Framework   string
}

func (r *Result) addLabel(name string, value string) {
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
		UUID:  generateUUID(),
		Start: getTimestampMs(),
	}
}
