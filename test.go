package allure

import (
	"fmt"
	"github.com/dailymotion/allure-go/severity"
	"github.com/dailymotion/allure-go/test"
	"github.com/fatih/camelcase"
	"github.com/jtolds/gls"
	"log"
	"runtime/debug"
	"strings"
	"testing"
)

type TestLabels struct {
	Epic        string
	Lead        string
	Owner       string
	Severity    severity.Severity
	Story       []string
	Feature     []string
	ParentSuite string
	Suite       string
	SubSuite    string
	Host        string
	Thread      string
	Framework   string
	Language    string
}

//Test execute the test and creates an Allure result used by Allure reports
func Test(t *testing.T, testOptions ...test.Option) {
	var r *Result
	r = newResult()
	r.UUID = generateUUID()
	r.Start = getTimestampMs()
	r.Name = t.Name()
	r.FullName = strings.Join(camelcase.Split(t.Name()), " ")
	r.Description = t.Name()
	r.setDefaultLabels(t)
	r.Steps = make([]StepObject, 0)
	for _, option := range testOptions {
		option(r)
	}

	defer func() {
		panicObject := recover()
		r.Stop = getTimestampMs()
		if panicObject != nil {
			t.Fail()
			r.StatusDetails = &StatusDetails{
				Message: fmt.Sprintf("%+v", panicObject),
				Trace:   filterStackTrace(debug.Stack()),
			}
			r.Status = broken
		}
		if r.Status == "" {
			r.Status = getTestStatus(t)
		}
		r.Stage = "finished"

		err := r.writeResultsFile()
		if err != nil {
			log.Println("Failed to write content of result to json file", err)
		}

		if panicObject != nil {
			panic(panicObject)
		}
	}()
	ctxMgr.SetValues(gls.Values{
		testResultKey:   r,
		nodeKey:         r,
		testInstanceKey: t,
	}, r.Test)
}
