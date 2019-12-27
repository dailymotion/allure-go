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
	StatusDetails *StatusDetails        `json:"statusDetails,omitempty"`
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

func (r *Result) AddReason(reason string) {
	testStatusDetails := r.StatusDetails
	if testStatusDetails == nil {
		testStatusDetails = &StatusDetails{}
	}
	r.StatusDetails.Message = reason
}

func (r *Result) AddDescription(description string) {
	r.Description = description
}

func (r *Result) AddParameter(name string, value interface{}) {
	r.Parameters = append(r.Parameters, parseParameter(name, value))
}

func (r *Result) AddName(name string) {
	r.Name = name
}

func (r *Result) AddAction(action func()) {
	r.Test = action
}

type FailureMode string

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
		UUID:  generateUUID(),
		Start: getTimestampMs(),
	}
}
