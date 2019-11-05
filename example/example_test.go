package example_test

import (
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/example"

	"testing"
)

func TestPassedExample(t *testing.T) {
	allure.Test(t, "This is a test to show allure implementation with a passing test", func() {
		s := "Hello world"
		if len(s) == 0 {
			t.Errorf("Expected 'hello world' string, but got %s ", s)
		}
	})
}

func TestWithIntricateSubsteps(t *testing.T) {
	allure.Test(t, "Test", func() {
		allure.Step("Step 1", func() {
			example.DoSomething()
			allure.Step("Sub-step 1.1", func() {
				t.Errorf("Failure")
			})
			allure.Step("Sub-step 1.2", func() {})
		})
		allure.Step("Step 2", func() {
			allure.Step("Sub-step 2.1", func() {
				allure.Step("Step 2.1.1", func() {
					allure.Step("Sub-step 2.1.1.1", func() {
						t.Errorf("Failure")
					})
					allure.Step("Sub-step 2.1.1.2", func() {
						t.Error("Failed like this")
					})
				})
			})
			allure.Step("Sub-step 2.2", func() {})
		})
	})
}

func TestFailedExample(t *testing.T) {
	allure.Test(t, "This is a test to show allure implementation with a failing test", func() {
		s := "Hello world"
		if len(s) != 0 {
			t.Errorf("Expected empty string, but got %s ", s)
		}
	})
}

func TestSkippedExample(t *testing.T) {
	allure.Test(t, "This is a test to show allure implementation with a skipped test", func() {
		t.Skip("Skipping this test")
	})
}
