package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestPassedExample(t *testing.T) {
	allure.Test(t,
		allure.Description("This is a test to show allure implementation with a passing test"),
		allure.Action(func() {
			s := "Hello world"
			if len(s) == 0 {
				t.Errorf("Expected 'hello world' string, but got %s ", s)
			}
		}))
}

func TestWithIntricateSubsteps(t *testing.T) {
	allure.Test(t, allure.Description("Test"),
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

func TestFailedExample(t *testing.T) {
	allure.Test(t, allure.Description("This is a test to show allure implementation with a failing test"), allure.Action(func() {
		s := "Hello world"
		if len(s) != 0 {
			t.Errorf("Expected empty string, but got %s ", s)
		}
	}))
}

func TestSkippedExample(t *testing.T) {
	allure.Test(t, allure.Description("This is a test to show allure implementation with a skipped test"), allure.Action(func() {
		t.Skip("Skipping this test")
	}))
}
