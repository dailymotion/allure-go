package allure

import (
	"github.com/jtolds/gls"
	"log"
)

type stepObject struct {
	Name          string       `json:"name,omitempty"`
	Status        string       `json:"status,omitempty"`
	Stage         string       `json:"stage"`
	ChildrenSteps []stepObject `json:"steps"`
	Attachements  []Attachment `json:"attachments"`
	Parameters    []Parameter  `json:"parameters"`
	Start         int64        `json:"start"`
	Stop          int64        `json:"stop"`
}

func (s *stepObject) GetSteps() []stepObject {
	return s.ChildrenSteps
}

func (s *stepObject) AddStep(step stepObject) {
	s.ChildrenSteps = append(s.ChildrenSteps, step)
}

// Step is meant to be wrapped around actions
func Step(description string, action func()) {
	StepWithParameter(description, nil, action)
}

// StepWithParameter is meant to be wrapped around actions with the purpose of logging the parameters
func StepWithParameter(description string, parameters map[string]interface{}, action func()) {
	step := newStep()
	step.Name = description
	step.Start = getTimestampMs()
	if parameters == nil || len(parameters) > 0 {
		step.Parameters = convertMapToParameters(parameters)
	}

	defer func() {
		step.Stop = getTimestampMs()
		currentStepObj, ok := ctxMgr.GetValue(nodeKey)
		if ok {
			currentStep := currentStepObj.(hasSteps)
			currentStep.AddStep(*step)
		} else {
			log.Fatalln("could not retrieve current node")
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
		Parameters:    make([]Parameter, 0),
	}
}
