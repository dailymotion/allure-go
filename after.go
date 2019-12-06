package allure

import (
	"github.com/jtolds/gls"
	"log"
	"testing"
)

// AfterTestWithParameters executes a teardown phase of the test and adds parameters to the Allure container object
func AfterTestWithParameters(t *testing.T, description string, parameters map[string]interface{}, labels TestLabels, testFunc func()) {
	after := newAfter()
	afterSubContainer := after.Afters[0]
	after.UUID = generateUUID()
	afterSubContainer.Start = getTimestampMs()
	after.Name = t.Name()

	after.Description = description

	afterSubContainer.Steps = make([]stepObject, 0)
	if parameters == nil || len(parameters) > 0 {
		afterSubContainer.Parameters = convertMapToParameters(parameters)
	}

	defer func() {
		afterSubContainer.Stop = getTimestampMs()
		afterSubContainer.Status = getTestStatus(t)
		afterSubContainer.Stage = "finished"
		testPhaseObject := getCurrentTestPhaseObject(t)
		if testPhaseObject.Test.UUID == "" {
			log.Printf("Test's \"%s\" allure teardwon is being executed before allure test!\n", t.Name())
		}

		after.Children = append(after.Children, testPhaseObject.Test.UUID)

		testPhaseObject.Afters = append(testPhaseObject.Afters, after)
		err := after.writeResultsFile()
		if err != nil {
			log.Println("Failed to write content of result to json file", err)
		}
	}()
	ctxMgr.SetValues(gls.Values{
		testResultKey:   afterSubContainer,
		nodeKey:         afterSubContainer,
		testInstanceKey: t,
	}, testFunc)
}

//AfterTest executes the teardown phase of the test and creates an Allure container object used by Allure reports
func AfterTest(t *testing.T, description string, testFunc func()) {
	AfterTestWithParameters(t, description, nil, TestLabels{}, testFunc)
}

func newAfter() *container {
	return &container{
		Children: make([]string, 0),
		Links:    make([]string, 0),
		Afters:   []*subContainer{newHelper()},
	}
}
