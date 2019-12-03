package allure

import (
	"testing"

	"github.com/jtolds/gls"
)

type stepObject struct {
	Name          string       `json:"name,omitempty"`
	Status        string       `json:"status,omitempty"`
	Stage         string       `json:"stage"`
	ChildrenSteps []stepObject `json:"steps"`
	Attachments   []attachment `json:"attachments"`
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

func (s *stepObject) GetAttachments() []attachment {
	return s.Attachments
}

func (s *stepObject) AddAttachment(attachment attachment) {
	s.Attachments = append(s.Attachments, attachment)
}

func (s *stepObject) SetStatus(status string) {
	s.Status = status
}

func (s *stepObject) GetStatus() string {
	return s.Status
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
			currentStep.AddStep(*step)
		})
	}()

	ctxMgr.SetValues(gls.Values{nodeKey: step}, action)
	step.Stage = "finished"
	if step.Status == "" {
		step.Status = "passed"
	}
}

func newStep() *stepObject {
	return &stepObject{
		Attachments:   make([]attachment, 0),
		ChildrenSteps: make([]stepObject, 0),
		Parameters:    make([]Parameter, 0),
	}
}
