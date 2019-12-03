package allure

import (
	"github.com/dailymotion/allure-go/severity"
	"github.com/fatih/camelcase"
	"github.com/jtolds/gls"
	"log"
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

// TestWithParameters executes a test and adds parameters to the Allure result object
func TestWithParameters(t *testing.T, description string, parameters map[string]interface{}, labels TestLabels, testFunc func()) {
	var r *result
	r = newResult()
	r.UUID = generateUUID()
	r.Start = getTimestampMs()
	r.Name = t.Name()
	r.FullName = strings.Join(camelcase.Split(t.Name()), " ")
	r.Description = description
	r.setLabels(t, labels)
	r.Steps = make([]stepObject, 0)
	if parameters == nil || len(parameters) > 0 {
		r.Parameters = convertMapToParameters(parameters)
	}

	defer func() {
		r.Stop = getTimestampMs()
		if r.Status == "" {
			r.Status = getTestStatus(t)
		}
		r.Stage = "finished"

		err := r.writeResultsFile()
		if err != nil {
			log.Println("Failed to write content of result to json file", err)
		}
	}()
	ctxMgr.SetValues(gls.Values{
		testResultKey:   r,
		nodeKey:         r,
		testInstanceKey: t,
	}, testFunc)
}

//Test execute the test and creates an Allure result used by Allure reports
func Test(t *testing.T, description string, testFunc func()) {
	TestWithParameters(t, description, nil, TestLabels{}, testFunc)
}
