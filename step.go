package allure

import "github.com/jtolds/gls"

type StepObject struct {
	Name          string `json:"name,omitempty"`
	Status        string `json:"status,omitempty"`
	Stage         string
	ChildrenSteps []StepObject `json:"steps"`
	Attachements  []Attachment `json:"attahcments"`
	Parameters    []string     `json:"parameters"`
	Start         int64
	Stop          int64
}

func Step(description string, testFunc func()) {
	ctxMgr.SetValues(gls.Values{}, testFunc)
}
