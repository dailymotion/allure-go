package allure

import (
	"testing"
)

const (
	broken  = "broken"
	passed  = "passed"
	failed  = "failed"
	skipped = "skipped"
)

type statusDetails struct {
	Known   bool   `json:"known,omitempty"`
	Muted   bool   `json:"muted,omitempty"`
	Flaky   bool   `json:"flaky,omitempty"`
	Message string `json:"message,omitempty"`
	Trace   string `json:"trace,omitempty"`
}

func getTestStatus(t *testing.T) string {
	if t.Failed() {
		return failed
	} else if t.Skipped() {
		return skipped
	} else {
		return passed
	}
}
