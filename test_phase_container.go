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
	Parameters    []Parameter    `json:"parameters,omitempty"`
}

func (sc *subContainer) GetAttachments() []attachment {
	return sc.Attachments
}

func (sc *subContainer) AddAttachment(attachment attachment) {
	sc.Attachments = append(sc.Attachments, attachment)
}

func (sc *subContainer) GetSteps() []stepObject {
	return sc.Steps
}

func (sc *subContainer) AddStep(step stepObject) {
	sc.Steps = append(sc.Steps, step)
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

//func getCurrentTestUUID(t *testing.T) string {
//	if currentTest := getCurrentTestPhaseObject(t).Test; currentTest != nil {
//		return currentTest.UUID
//	} else {
//		return ""
//	}
//}

func (c container) writeResultsFile() error {
	createFolderOnce.Do(createFolderIfNotExists)
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
		Parameters:  make([]Parameter, 0),
	}
}
