package allure

import (
	"fmt"
	"github.com/jtolds/gls"
	"log"
	"os"
	"testing"
)

// TestWithParameters executes a test and adds parameters to the Allure result object
func TestWithParameters(t *testing.T, description string, parameters map[string]interface{}, testFunc func()) {
	var r *result
	r = newResult()
	r.UUID = generateUUID()
	r.Start = getTimestampMs()
	r.Name = t.Name()
	r.Description = description
	r.setLabels(t)
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
	ctxMgr.SetValues(gls.Values{
		testResultKey:   r,
		nodeKey:         r,
		testInstanceKey: t,
	}, testFunc)
}

//Test execute the test and creates an Allure result used by Allure reports
func Test(t *testing.T, description string, testFunc func()) {
	TestWithParameters(t, description, nil, testFunc)
}
