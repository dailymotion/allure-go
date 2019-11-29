package allure

import (
	"log"
	"testing"

	"github.com/jtolds/gls"
)

type stepObject struct {
	Name          string        `json:"name,omitempty"`
	Status        string        `json:"status,omitempty"`
	StatusDetail  statusDetails `json:"statusDetails,omitempty"`
	Stage         string        `json:"stage"`
	ChildrenSteps []stepObject  `json:"steps"`
	Attachments   []attachment  `json:"attachments"`
	Parameters    []Parameter   `json:"parameters"`
	Start         int64         `json:"start"`
	Stop          int64         `json:"stop"`
}

func (s *stepObject) GetSteps() []stepObject {
	return s.ChildrenSteps
}

func (s *stepObject) AddStep(step stepObject) {
	s.ChildrenSteps = append(s.ChildrenSteps, step)
}

func (s *stepObject) GetAttachments() []attachment {
	return s.Attachments
}

func (s *stepObject) AddAttachment(attachment attachment) {
	s.Attachments = append(s.Attachments, attachment)
}

// Step is meant to be wrapped around actions
func Step(description string, action func()) {
	StepWithParameter(description, nil, action)
}

// SkipStep doesn't execute the action and marks the step as skipped in report
// Reason won't appear in report until https://github.com/allure-framework/allure2/issues/774 is fixed
func SkipStep(description, reason string, action func()) {
	SkipStepWithParameter(description, reason, nil, action)
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
		if testInstance, ok := ctxMgr.GetValue(testInstanceKey); ok {
			if testInstance.(*testing.T).Failed() {
				step.Status = "failed"
			}
		}

		if currentStepObj, ok := ctxMgr.GetValue(nodeKey); ok {
			currentStep := currentStepObj.(hasSteps)
			currentStep.AddStep(*step)
		} else {
			log.Fatalln("could not retrieve current allure node")
		}
	}()

	ctxMgr.SetValues(gls.Values{nodeKey: step}, action)
	step.Stage = "finished"
	step.Status = "passed"
}

// SkipStepWithParameter doesn't execute the action and marks the step as skipped in report
// Reason won't appear in report until https://github.com/allure-framework/allure2/issues/774 is fixed
func SkipStepWithParameter(description, reason string, parameters map[string]interface{}, action func()) {
	step := newStep()
	step.Start = getTimestampMs()
	step.Name = description
	if parameters == nil || len(parameters) > 0 {
		step.Parameters = convertMapToParameters(parameters)
	}
	step.Status = "skipped"
	step.StatusDetail.Message = reason
	if currentStepObj, ok := ctxMgr.GetValue(nodeKey); ok {
		currentStep := currentStepObj.(hasSteps)
		currentStep.AddStep(*step)
	} else {
		log.Fatalln("could not retrieve current allure node")
	}
	step.Stop = getTimestampMs()
}

func newStep() *stepObject {
	return &stepObject{
		Attachments:   make([]attachment, 0),
		ChildrenSteps: make([]stepObject, 0),
		Parameters:    make([]Parameter, 0),
	}
}
