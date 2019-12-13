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

func (s *stepObject) getSteps() []stepObject {
	return s.ChildrenSteps
}

func (s *stepObject) addStep(step stepObject) {
	s.ChildrenSteps = append(s.ChildrenSteps, step)
}

func (s *stepObject) getAttachments() []attachment {
	return s.Attachments
}

func (s *stepObject) addAttachment(attachment attachment) {
	s.Attachments = append(s.Attachments, attachment)
}

func (s *stepObject) setStatus(status string) {
	s.Status = status
}

func (s *stepObject) getStatus() string {
	return s.Status
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
		manipulateOnObjectFromCtx(
			testInstanceKey,
			func(testInstance interface{}) {
				if testInstance.(*testing.T).Failed() {
					if step.Status == "" {
						step.Status = "failed"
					}
				}
			})
		manipulateOnObjectFromCtx(nodeKey, func(currentStepObj interface{}) {
			currentStep := currentStepObj.(hasSteps)
			currentStep.addStep(*step)
		})
	}()

	ctxMgr.SetValues(gls.Values{nodeKey: step}, action)
	step.Stage = "finished"
	if step.Status == "" {
		step.Status = "passed"
	}
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
		currentStep.addStep(*step)
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
