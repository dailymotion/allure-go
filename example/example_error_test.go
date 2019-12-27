package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/step"
	"github.com/dailymotion/allure-go/test"
	"github.com/pkg/errors"
	"testing"
)

func TestFailAllure(t *testing.T) {
	allure.Test(t, test.Description("Test with Allure error in it"), test.Body(func() {
		allure.Step(step.Description("Step 1"), step.Action(func() {
		}))
		allure.Step(step.Description("Step 2"), step.Action(func() {
			allure.Fail(errors.New("Error message"))
		}))
	}))
}

func TestFailNowAllure(t *testing.T) {
	allure.Test(t, test.Description("Test with Allure error in it"), test.Body(func() {
		allure.FailNow(errors.New("A more serious error"))
		allure.Step(step.Description("Step you're not supposed to see"), step.Action(func() {}))
	}))
}

func TestBreakAllure(t *testing.T) {
	allure.Test(t, test.Description("Test with Allure error in it"), test.Body(func() {
		allure.Step(step.Description("Step 1"), step.Action(func() {}))
		allure.Step(step.Description("Step 2"), step.Action(func() {
			allure.Break(errors.New("Error message"))
		}))
	}))
}

func TestBreakNowAllure(t *testing.T) {
	allure.Test(t, test.Description("Test with Allure error in it"), test.Body(func() {
		allure.BreakNow(errors.New("A more serious error"))
		allure.Step(step.Description("Step you're not supposed to see"), step.Action(func() {}))
	}))
}
