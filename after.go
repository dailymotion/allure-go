package allure

import (
	"fmt"
	"github.com/jtolds/gls"
	"log"
	"os"
	"testing"
)

// TestWithParameters executes a test and adds parameters to the Allure result object
func AfterWithParameters(t *testing.T, description string, parameters map[string]interface{}, labels TestLabels, testFunc func()) {
	after := newAfter()
	afterHelper := newAfterHelper()

	after.UUID = generateUUID()
	afterHelper.Start = getTimestampMs()
	after.Name = t.Name()

	after.Description = description

	afterHelper.Steps = make([]stepObject, 0)
	if parameters == nil || len(parameters) > 0 {
		afterHelper.Parameters = convertMapToParameters(parameters)
	}

	defer func() {
		afterHelper.Stop = getTimestampMs()
		afterHelper.Status = getTestStatus(t)
		afterHelper.Stage = "finished"

		err := after.writeResultsFile()
		if err != nil {
			log.Fatalf(fmt.Sprintf("Failed to write content of result to json file"), err)
			os.Exit(1)
		}
	}()
	ctxMgr.SetValues(gls.Values{
		testResultKey:   afterHelper,
		nodeKey:         afterHelper,
		testInstanceKey: t,
	}, testFunc)
}

//Test execute the test and creates an Allure result used by Allure reports
func After(t *testing.T, description string, testFunc func()) {
	TestWithParameters(t, description, nil, TestLabels{}, testFunc)
}

func newAfter() *Container {
	return &Container{
		Afters: []helperContainer{*(newAfterHelper())},
	}
}

func newAfterHelper() *helperContainer {
	return newHelper()
}
