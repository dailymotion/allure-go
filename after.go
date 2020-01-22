package allure

import (
	"github.com/jtolds/gls"
	"log"
	"testing"
)

// AfterTest executes a teardown phase of the test and adds parameters to the Allure container object
func AfterTest(t *testing.T, testOptions ...Option) {
	after := newAfter()
	afterSubContainer := after.Afters[0]
	after.UUID = generateUUID()
	afterSubContainer.Start = getTimestampMs()

	afterSubContainer.Steps = make([]stepObject, 0)
	for _, option := range testOptions {
		option(afterSubContainer)
	}

	defer func() {
		afterSubContainer.Stop = getTimestampMs()
		afterSubContainer.Status = getTestStatus(t)
		afterSubContainer.Stage = "finished"
		testPhaseObject := getCurrentTestPhaseObject(t)
		if testPhaseObject.Test.UUID == "" {
			log.Printf("Test's \"%s\" allure teardщwn is being executed before allure test!\n", t.Name())
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
	}, afterSubContainer.Action)
}

func newAfter() *container {
	return &container{
		Children: make([]string, 0),
		Links:    make([]string, 0),
		Afters:   []*subContainer{newHelper()},
	}
}
