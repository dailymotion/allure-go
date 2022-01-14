package allure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

//result is the top level report object for a test
type result struct {
	UUID          string         `json:"uuid,omitempty"`
	TestCaseID    string         `json:"testCaseId,omitempty"`
	HistoryID     string         `json:"historyId,omitempty"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetails *statusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage,omitempty"`
	Steps         []stepObject   `json:"steps,omitempty"`
	Attachments   []attachment   `json:"attachments,omitempty"`
	Parameters    []parameter    `json:"parameters,omitempty"`
	Start         int64          `json:"start,omitempty"`
	Stop          int64          `json:"stop,omitempty"`
	Children      []string       `json:"children,omitempty"`
	FullName      string         `json:"fullName,omitempty"`
	Labels        []label        `json:"labels,omitempty"`
	Test          func()         `json:"-"`
}

func (r *result) addReason(reason string) {
	testStatusDetails := r.StatusDetails
	if testStatusDetails == nil {
		r.StatusDetails = &statusDetails{}
	}
	r.StatusDetails.Message = reason
}

func (r *result) addDescription(description string) {
	r.Description = description
}

func (r *result) addParameter(name string, value interface{}) {
	r.Parameters = append(r.Parameters, parseParameter(name, value))
}

func (r *result) addParameters(parameters map[string]interface{}) {
	for key, value := range parameters {
		r.Parameters = append(r.Parameters, parseParameter(key, value))
	}
}

func (r *result) addName(name string) {
	r.Name = name
}

func (r *result) addAction(action func()) {
	r.Test = action
}

type FailureMode string

func (r *result) getAttachments() []attachment {
	return r.Attachments
}

func (r *result) addAttachment(attachment attachment) {
	r.Attachments = append(r.Attachments, attachment)
}

func (r *result) getSteps() []stepObject {
	return r.Steps
}

func (r *result) addStep(step stepObject) {
	r.Steps = append(r.Steps, step)
}

func (r *result) addFullName(FullName string) {
	r.FullName = FullName
}

func (r *result) setStatus(status string) {
	r.Status = status
}

func (r *result) getStatus() string {
	return r.Status
}

func (r *result) getStatusDetails() *statusDetails {
	return r.StatusDetails
}

func (r *result) setStatusDetails(details statusDetails) {
	r.StatusDetails = &details
}

func (r *result) setDefaultLabels(t *testing.T) {
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

	r.addLabel("package", testPackage)
	r.addLabel("testClass", testPackage)
	r.addLabel("testMethod", t.Name())
	if len(wsd) == 0 {
		r.addFullName(fmt.Sprintf("%s:%s", testFile, t.Name()))
	} else {
		r.addFullName(fmt.Sprintf("%s:%s", strings.TrimPrefix(testFile, wsd+"/"), t.Name()))
	}
	if hostname, err := os.Hostname(); err == nil {
		r.addLabel("host", hostname)
	}

	r.addLabel("language", "golang")

	//TODO: these labels are available, but should be handled separately.

	//	Thread      string
	//	Framework   string
}

func (r *result) addLabel(name string, value string) {
	r.Labels = append(r.Labels, label{
		Name:  name,
		Value: value,
	})
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
