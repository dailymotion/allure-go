package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/step"
	"github.com/dailymotion/allure-go/test"

	"testing"
)

func TestPassedExample(t *testing.T) {
	allure.Test(t,
		test.Description("This is a test to show allure implementation with a passing test"),
		test.Body(func() {
			s := "Hello world"
			if len(s) == 0 {
				t.Errorf("Expected 'hello world' string, but got %s ", s)
			}
		}))
}

func TestWithIntricateSubsteps(t *testing.T) {
	allure.Test(t, test.Description("Test"),
		test.Body(func() {
			allure.Step(step.Description("Step 1"), step.Action(func() {
				doSomething()
				allure.Step(step.Description("Sub-step 1.1"), step.Action(func() {
					t.Errorf("Failure")
				}))
				allure.Step(step.Description("Sub-step 1.2"), step.Action(func() {}))
				allure.SkipStep(step.Description("Sub-step 1.3"), step.Reason("Skip this step because of defect to be fixed"), step.Action(func() {}))
			}))
			allure.Step(step.Description("Step 20"), step.Action(func() {
				allure.Step(step.Description("Sub-step 2.1"), step.Action(func() {
					allure.Step(step.Description("Step 2.1.1"), step.Action(func() {
						allure.Step(step.Description("Sub-step 2.1.1.1"), step.Action(func() {
							t.Errorf("Failure")
						}))
						allure.Step(step.Description("Sub-step 2.1.1.2"), step.Action(func() {
							t.Error("Failed like this")
						}))
					}))
				}))
				allure.Step(step.Description("Sub-step 2.2"), step.Action(func() {}))
			}))
		}))
}

func TestFailedExample(t *testing.T) {
	allure.Test(t, test.Description("This is a test to show allure implementation with a failing test"), test.Body(func() {
		s := "Hello world"
		if len(s) != 0 {
			t.Errorf("Expected empty string, but got %s ", s)
		}
	}))
}

func TestSkippedExample(t *testing.T) {
	allure.Test(t, test.Description("This is a test to show allure implementation with a skipped test"), test.Body(func() {
		t.Skip("Skipping this test")
	}))
}
