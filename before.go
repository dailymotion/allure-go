package allure

import (
	"fmt"
	"github.com/fatih/camelcase"
	"github.com/jtolds/gls"
	"log"
	"runtime/debug"
	"strings"
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
		panicObject := recover()

		beforeSubContainer.Stop = getTimestampMs()
		beforeSubContainer.Status = getTestStatus(t)
		beforeSubContainer.Stage = "finished"

		if panicObject != nil {
			t.Fail()
			beforeSubContainer.StatusDetails = &statusDetails{
				Message: fmt.Sprintf("%+v", panicObject),
				Trace:   filterStackTrace(debug.Stack()),
			}
			beforeSubContainer.Status = broken

			r := newResult()
			r.Stop = getTimestampMs()
			r.Name = strings.Join(camelcase.Split(t.Name())[1:], " ")
			r.Description = t.Name()
			r.setDefaultLabels(t)
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
			panic(panicObject)
		}
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
