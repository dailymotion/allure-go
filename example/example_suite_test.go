package example

import (
	"testing"

	"github.com/dailymotion/allure-go"
)

func TestSuiteExample(t *testing.T) {
	allure.Test(t,
		allure.ParentSuite("Epic Root Suite"),
		allure.Suite("Test Suite"),
		allure.SubSuite("Sub Suite Example"),
		allure.Description("This is a test to show allure implementation with a passing test"),
		allure.Action(func() {
			s := "Hello world"
			if len(s) == 0 {
				t.Errorf("Expected 'hello world' string, but got %s ", s)
			}
		}))
}

func TestSuiteWithIntricateSubsteps(t *testing.T) {
	allure.Test(t,
		allure.ParentSuite("Epic Root Suite"),
		allure.Suite("Test Suite"),
		allure.SubSuite("Sub Suite Example"),
		allure.Description("Test"),
		allure.Action(func() {
			allure.Step(allure.Description("Step 1"), allure.Action(func() {
				doSomething()
				allure.Step(allure.Description("Sub-step 1.1"), allure.Action(func() {
					t.Errorf("Failure")
				}))
				allure.Step(allure.Description("Sub-step 1.2"), allure.Action(func() {}))
				allure.SkipStep(allure.Description("Sub-step 1.3"), allure.Reason("Skip this step because of defect to be fixed"), allure.Action(func() {}))
			}))
			allure.Step(allure.Description("Step 2"), allure.Action(func() {
				allure.Step(allure.Description("Sub-step 2.1"), allure.Action(func() {
					allure.Step(allure.Description("Step 2.1.1"), allure.Action(func() {
						allure.Step(allure.Description("Sub-step 2.1.1.1"), allure.Action(func() {
							t.Errorf("Failure")
						}))
						allure.Step(allure.Description("Sub-step 2.1.1.2"), allure.Action(func() {
							t.Error("Failed like this")
						}))
					}))
				}))
				allure.Step(allure.Description("Sub-step 2.2"), allure.Action(func() {}))
			}))
		}))
}
