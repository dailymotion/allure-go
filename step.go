package allure

import (
	"github.com/dailymotion/allure-go/parameter"
	"github.com/dailymotion/allure-go/step"
	"github.com/pkg/errors"
	"log"
	"testing"

	"github.com/jtolds/gls"
)

type StepObject struct {
	Name          string                `json:"name,omitempty"`
	Status        string                `json:"status,omitempty"`
	StatusDetails *StatusDetails        `json:"statusDetails,omitempty"`
	Stage         string                `json:"stage"`
	ChildrenSteps []StepObject          `json:"steps"`
	Attachments   []Attachment          `json:"attachments"`
	Parameters    []parameter.Parameter `json:"parameters"`
	Start         int64                 `json:"start"`
	Stop          int64                 `json:"stop"`
	Action        func()                `json:"-"`
}

func (s *StepObject) AddReason(reason string) {
	testStatusDetails := s.StatusDetails
	if testStatusDetails == nil {
		testStatusDetails = &StatusDetails{}
	}
	s.StatusDetails.Message = reason
}

func (s *StepObject) AddLabel(key string, value string) {
	// Step doesn't have labels
}

func (s *StepObject) AddDescription(description string) {
	// Step doesn't have description
}

func (s *StepObject) AddParameter(name string, value interface{}) {
	s.Parameters = append(s.Parameters, parseParameter(name, value))
}

func (s *StepObject) AddName(name string) {
	s.Name = name
}

func (s *StepObject) AddAction(action func()) {
	s.Action = action
}

func (s *StepObject) GetSteps() []StepObject {
	return s.ChildrenSteps
}

func (s *StepObject) AddStep(step StepObject) {
	s.ChildrenSteps = append(s.ChildrenSteps, step)
}

func (s *StepObject) GetAttachments() []Attachment {
	return s.Attachments
}

func (s *StepObject) AddAttachment(attachment Attachment) {
	s.Attachments = append(s.Attachments, attachment)
}

func (s *StepObject) SetStatus(status string) {
	s.Status = status
}

func (s *StepObject) GetStatus() string {
	return s.Status
}

// SkipStep doesn't execute the action and marks the step as skipped in report
// Reason won't appear in report until https://github.com/allure-framework/allure2/issues/774 is fixed
func SkipStep(stepOptions ...step.Option) {
	stepObject := newStep()
	stepObject.Start = getTimestampMs()
	for _, option := range stepOptions {
		option(stepObject)
	}
	stepObject.Status = "skipped"
	if currentStepObj, ok := ctxMgr.GetValue(nodeKey); ok {
		currentStep := currentStepObj.(hasSteps)
		currentStep.AddStep(*stepObject)
	} else {
		log.Fatalln("could not retrieve current allure node")
	}
	stepObject.Stop = getTimestampMs()
}

// Step is meant to be wrapped around actions
func Step(stepOptions ...step.Option) {
	stepObject := newStep()
	stepObject.Start = getTimestampMs()
	for _, option := range stepOptions {
		option(stepObject)
	}

	defer func() {
		panicObject := recover()
		stepObject.Stop = getTimestampMs()
		manipulateOnObjectFromCtx(
			testInstanceKey,
			func(testInstance interface{}) {
				if panicObject != nil {
					Break(errors.Errorf("%+v", panicObject))
				}
				if testInstance.(*testing.T).Failed() ||
					panicObject != nil {
					if stepObject.Status == "" {
						stepObject.Status = "failed"
					}
				}
			})
		stepObject.Stage = "finished"
		if stepObject.Status == "" {
			stepObject.Status = "passed"
		}
		manipulateOnObjectFromCtx(nodeKey, func(currentStepObj interface{}) {
			currentStep := currentStepObj.(hasSteps)
			currentStep.AddStep(*stepObject)
		})

		if panicObject != nil {
			panic(panicObject)
		}
	}()

	ctxMgr.SetValues(gls.Values{nodeKey: stepObject}, stepObject.Action)
}

func newStep() *StepObject {
	return &StepObject{
		Attachments:   make([]Attachment, 0),
		ChildrenSteps: make([]StepObject, 0),
		Parameters:    make([]parameter.Parameter, 0),
	}
}
