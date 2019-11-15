package allure

import (
	"fmt"
	"github.com/fatih/camelcase"
	"github.com/jtolds/gls"
	"log"
	"os"
	"strings"
	"testing"
)

type TestLabels struct {
	Epic        string
	Lead        string
	Owner       string
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
		r.Status = getTestStatus(t)
		r.Stage = "finished"

		err := r.writeResultsFile()
		if err != nil {
			log.Fatalf(fmt.Sprintf("Failed to write content of result to json file"), err)
			os.Exit(1)
		}
	}()
	ctxMgr.SetValues(gls.Values{"test_result_object": r, nodeKey: r}, testFunc)
}

//Test execute the test and creates an Allure result used by Allure reports
func Test(t *testing.T, description string, testFunc func()) {
	TestWithParameters(t, description, nil, TestLabels{}, testFunc)
}
