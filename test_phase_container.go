package allure

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"testing"
)

type testPhaseContainer struct {
	Befores []*container
	Test    *result
	Afters  []*container
}

type container struct {
	UUID        string          `json:"uuid"`
	Name        string          `json:"name"`
	Children    []string        `json:"children"`
	Description string          `json:"description"`
	Befores     []*subContainer `json:"befores"`
	Afters      []*subContainer `json:"afters"`
	Links       []string        `json:"links"`
	Start       int64           `json:"start"`
	Stop        int64           `json:"stop"`
}

//subContainer defines a step
type subContainer struct {
	Name          string         `json:"name,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetails *statusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage,omitempty"`
	Description   string         `json:"description,omitempty"`
	Start         int64          `json:"start,omitempty"`
	Stop          int64          `json:"stop,omitempty"`
	Steps         []stepObject   `json:"steps,omitempty"`
	Attachments   []attachment   `json:"attachments,omitempty"`
	Parameters    []parameter    `json:"parameters,omitempty"`
	Action        func()         `json:"-"`
}

func (sc *subContainer) addLabel(key string, value string) {
	panic("implement me")
}

func (sc *subContainer) addDescription(description string) {
	sc.Description = description
}

func (sc *subContainer) addParameter(name string, value interface{}) {
	sc.Parameters = append(sc.Parameters, parseParameter(name, value))
}

func (sc *subContainer) addParameters(parameters map[string]interface{}) {
	for key, value := range parameters {
		sc.Parameters = append(sc.Parameters, parseParameter(key, value))
	}
}

func (sc *subContainer) addName(name string) {
	sc.Name = name
}

func (sc *subContainer) addAction(action func()) {
	sc.Action = action
}

func (sc *subContainer) addReason(reason string) {
	testStatusDetails := sc.StatusDetails
	if testStatusDetails == nil {
		testStatusDetails = &statusDetails{}
	}
	sc.StatusDetails.Message = reason
}

func (sc *subContainer) getAttachments() []attachment {
	return sc.Attachments
}

func (sc *subContainer) addAttachment(attachment attachment) {
	sc.Attachments = append(sc.Attachments, attachment)
}

func (sc *subContainer) getSteps() []stepObject {
	return sc.Steps
}

func (sc *subContainer) addStep(step stepObject) {
	sc.Steps = append(sc.Steps, step)
}

func (sc *subContainer) getStatusDetails() *statusDetails {
	return sc.StatusDetails
}

func (sc *subContainer) setStatusDetails(details statusDetails) {
	sc.StatusDetails = &details
}

func (sc *subContainer) setStatus(status string) {
	sc.Status = status
}

func (sc *subContainer) getStatus() string {
	return sc.Status
}

func getCurrentTestPhaseObject(t *testing.T) *testPhaseContainer {
	var currentPhaseObject *testPhaseContainer
	if phaseContainer, ok := testPhaseObjects[t.Name()]; ok {
		currentPhaseObject = phaseContainer
	} else {
		currentPhaseObject = &testPhaseContainer{
			Befores: make([]*container, 0),
			Afters:  make([]*container, 0),
		}
		testPhaseObjects[t.Name()] = currentPhaseObject
	}

	return currentPhaseObject
}

func (c container) writeResultsFile() error {
	ensureFolderCreated()
	copyEnvFileOnce.Do(copyEnvFileIfExists)

	j, err := json.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "Failed to marshall result into JSON")
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s-container.json", resultsPath, c.UUID), j, 0777)
	if err != nil {
		return errors.Wrap(err, "Failed to write in file")
	}
	return nil
}

func newHelper() *subContainer {
	return &subContainer{
		Steps:       make([]stepObject, 0),
		Attachments: make([]attachment, 0),
		Parameters:  make([]parameter, 0),
	}
}
