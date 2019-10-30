package allure

import (
	"fmt"
	"github.com/jtolds/gls"
	"log"
	"time"
)

type stepObject struct {
	Name          string       `json:"name,omitempty"`
	Status        string       `json:"status,omitempty"`
	Stage         string       `json:"stage"`
	ChildrenSteps []stepObject `json:"steps"`
	Attachements  []Attachment `json:"attachments"`
	Parameters    []string     `json:"parameters"`
	Start         int64        `json:"start"`
	Stop          int64        `json:"stop"`
}

//"status": "passed",
//      "stage": "finished",
//      "steps": [
//        {
//          "name": "Other stuff",
//          "status": "passed",
//          "stage": "finished",
//          "steps": [],
//          "attachments": [],
//          "parameters": [],
//          "start": 1572368980980,
//          "stop": 1572368980981
//        }
//      ],
//      "attachments": [],
//      "parameters": [],
//      "start": 1572368980979,
//      "stop": 1572368980982

func (s *stepObject) GetSteps() []stepObject {
	return s.ChildrenSteps
}

func (s *stepObject) AddStep(step stepObject) {
	s.ChildrenSteps = append(s.ChildrenSteps, step)
}

func Step(description string, action func()) {
	log.Println(description)
	step := newStep()
	step.Name = description
	step.Start = time.Now().Unix()
	defer func() {
		step.Stop = time.Now().Unix()
		currentStepObj, ok := ctxMgr.GetValue(nodeKey)
		if ok {
			currentStep := currentStepObj.(HasSteps)
			currentStep.AddStep(*step)
		} else {
			fmt.Println("WTF?")
		}

	}()

	ctxMgr.SetValues(gls.Values{nodeKey: step}, action)
	step.Stage = "finished"
	step.Status = "passed"
}

func newStep() *stepObject {
	return &stepObject{
		Attachements:  make([]Attachment, 0),
		ChildrenSteps: make([]stepObject, 0),
	}
}
