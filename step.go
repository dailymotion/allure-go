package allure

import (
	"log"
	"testing"

	"github.com/pkg/errors"

	"github.com/jtolds/gls"
)

type stepObject struct {
	Name          string         `json:"name,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetails *statusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage"`
	ChildrenSteps []stepObject   `json:"steps"`
	Attachments   []attachment   `json:"attachments"`
	Parameters    []parameter    `json:"parameters"`
	Start         int64          `json:"start"`
	Stop          int64          `json:"stop"`
	Action        func()         `json:"-"`
}

func (s *stepObject) addReason(reason string) {
	testStatusDetails := s.StatusDetails
	if testStatusDetails == nil {
		s.StatusDetails = &statusDetails{}
	}
	s.StatusDetails.Message = reason
}

func (s *stepObject) addLabel(key string, value string) {
	// Step doesn't have labels
}

func (s *stepObject) addDescription(description string) {
	s.Name = description
}

func (s *stepObject) addParameter(name string, value interface{}) {
	s.Parameters = append(s.Parameters, parseParameter(name, value))
}

func (s *stepObject) addParameters(parameters map[string]interface{}) {
	for key, value := range parameters {
		s.Parameters = append(s.Parameters, parseParameter(key, value))
	}
}

func (s *stepObject) addName(name string) {
	s.Name = name
}

func (s *stepObject) addAction(action func()) {
	s.Action = action
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

// SkipStep doesn't execute the action and marks the step as skipped in report
// Reason won't appear in report until https://github.com/allure-framework/allure2/issues/774 is fixed
func SkipStep(stepOptions ...Option) {
	stepObject := newStep()
	stepObject.Start = getTimestampMs()
	for _, option := range stepOptions {
		option(stepObject)
	}
	stepObject.Status = "skipped"
	stepObject.Stage = "finished"
	stepObject.Stop = getTimestampMs()
	if currentStepObj, ok := ctxMgr.GetValue(nodeKey); ok {
		currentStep := currentStepObj.(hasSteps)
		currentStep.addStep(*stepObject)
	} else {
		log.Fatalln("could not retrieve current allure node")
	}
}

// Step is meant to be wrapped around actions
func Step(stepOptions ...Option) {
	stepObject := newStep()
	stepObject.Start = getTimestampMs()
	for _, option := range stepOptions {
		option(stepObject)
	}
	if stepObject.Action == nil {
		stepObject.Action = func() {}
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
			currentStep.addStep(*stepObject)
		})

		if panicObject != nil {
			panic(panicObject)
		}
	}()

	ctxMgr.SetValues(gls.Values{nodeKey: stepObject}, stepObject.Action)
}

func newStep() *stepObject {
	return &stepObject{
		Attachments:   make([]attachment, 0),
		ChildrenSteps: make([]stepObject, 0),
		Parameters:    make([]parameter, 0),
	}
}
