package allure

import (
	"github.com/jtolds/gls"
	"log"
	"testing"
)

// BeforeTestWithParameters executes a setup phase of the test and adds parameters to the Allure container object
func BeforeTestWithParameters(t *testing.T, description string, parameters map[string]interface{}, labels TestLabels, testFunc func()) {
	testPhaseObject := getCurrentTestPhaseObject(t)
	if testPhaseObject.Test != nil {
		log.Printf("Test's \"%s\" allure setup is being executed after allure test!\n", t.Name())
	}
	before := newBefore()
	testPhaseObject.Befores = append(testPhaseObject.Befores, before)
	beforeSubContainer := before.Befores[0]

	before.UUID = generateUUID()
	beforeSubContainer.Start = getTimestampMs()
	before.Name = t.Name()

	before.Description = description

	beforeSubContainer.Steps = make([]stepObject, 0)
	if parameters == nil || len(parameters) > 0 {
		beforeSubContainer.Parameters = convertMapToParameters(parameters)
	}

	defer func() {
		beforeSubContainer.Stop = getTimestampMs()
		beforeSubContainer.Status = getTestStatus(t)
		beforeSubContainer.Stage = "finished"

	}()
	ctxMgr.SetValues(gls.Values{
		testResultKey:   beforeSubContainer,
		nodeKey:         beforeSubContainer,
		testInstanceKey: t,
	}, testFunc)
}

//BeforeTest executes the setup phase of the test and creates an Allure container object used by Allure reports
func BeforeTest(t *testing.T, description string, testFunc func()) {
	BeforeTestWithParameters(t, description, nil, TestLabels{}, testFunc)
}

func newBefore() *container {
	return &container{
		Befores: []*subContainer{newHelper()},
	}
}
