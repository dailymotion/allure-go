package allure

import (
	"fmt"
	"github.com/dailymotion/allure-go/severity"
	"github.com/fatih/camelcase"
	"github.com/jtolds/gls"
	"log"
	"runtime/debug"
	"strings"
	"testing"
)

type testLabels struct {
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

func SkipTest(t *testing.T, testOptions ...Option) {
	var r *result
	r = newResult()
	r.UUID = generateUUID()
	r.Start = getTimestampMs()
	r.Name = strings.Join(camelcase.Split(t.Name())[1:], " ")
	r.Description = t.Name()
	r.setDefaultLabels(t)
	r.Steps = make([]stepObject, 0)
	for _, option := range testOptions {
		option(r)
	}

	getCurrentTestPhaseObject(t).Test = r
	r.Stop = getTimestampMs()
	r.Stage = "finished"
	r.Status = skipped
	err := r.writeResultsFile()
	if err != nil {
		log.Println("Failed to write content of result to json file", err)
	}
	setups := getCurrentTestPhaseObject(t).Befores
	for _, setup := range setups {
		setup.Children = append(setup.Children, r.UUID)
		err := setup.writeResultsFile()
		if err != nil {
			log.Println("Failed to write content of result to json file", err)
		}
	}
	defer func() {
		if r.StatusDetails != nil {
			t.Skip(r.StatusDetails.Message)
		} else {
			t.Skip()
		}
	}()
}

//Test execute the test and creates an Allure result used by Allure reports
func Test(t *testing.T, testOptions ...Option) {
	var r *result
	r = newResult()
	r.UUID = generateUUID()
	r.Start = getTimestampMs()
	r.Name = strings.Join(camelcase.Split(t.Name())[1:], " ")
	r.Description = t.Name()
	r.setDefaultLabels(t)
	r.Steps = make([]stepObject, 0)
	for _, option := range testOptions {
		option(r)
	}

	if r.Test == nil {
		r.Test = func() {}
	}

	defer func() {
		panicObject := recover()
		getCurrentTestPhaseObject(t).Test = r
		r.Stop = getTimestampMs()
		if panicObject != nil {
			isFailed := t.Failed()
			t.Fail()
			r.StatusDetails = &statusDetails{
				Message: fmt.Sprintf("%+v", panicObject),
				Trace:   filterStackTrace(debug.Stack()),
			}
			if !isFailed {
				r.Status = broken
			}
		}
		if r.Status == "" {
			r.Status = getTestStatus(t)
		}
		r.Stage = "finished"

		err := r.writeResultsFile()
		if err != nil {
			log.Println("Failed to write content of result to json file", err)
		}
		setups := getCurrentTestPhaseObject(t).Befores
		for _, setup := range setups {
			setup.Children = append(setup.Children, r.UUID)
			err := setup.writeResultsFile()
			if err != nil {
				log.Println("Failed to write content of result to json file", err)
			}
		}

		//if panicObject != nil {
		//	panic(panicObject)
		//}
	}()
	ctxMgr.SetValues(gls.Values{
		testResultKey:   r,
		nodeKey:         r,
		testInstanceKey: t,
	}, r.Test)
}
