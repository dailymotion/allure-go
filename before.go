package allure

import (
	"github.com/jtolds/gls"
	"log"
	"testing"
)

// BeforeTest executes a setup phase of the test and adds parameters to the Allure container object
func BeforeTest(t *testing.T, testOptions ...Option) {
	testPhaseObject := getCurrentTestPhaseObject(t)
	if testPhaseObject.Test != nil {
		log.Printf("Test's \"%s\" allure setup is being executed after allure test!\n", t.Name())
	}
	before := newBefore()
	testPhaseObject.Befores = append(testPhaseObject.Befores, before)
	beforeSubContainer := before.Befores[0]

	before.UUID = generateUUID()
	beforeSubContainer.Start = getTimestampMs()
	beforeSubContainer.Steps = make([]stepObject, 0)
	for _, option := range testOptions {
		option(beforeSubContainer)
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
	}, beforeSubContainer.Action)
}

func newBefore() *container {
	return &container{
		Befores: []*subContainer{newHelper()},
	}
}
