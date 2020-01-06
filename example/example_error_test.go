package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/pkg/errors"
	"testing"
)

func TestFailAllure(t *testing.T) {
	allure.Test(t, allure.Description("Test with Allure error in it"), allure.Body(func() {
		allure.Step(allure.Description("Step 1"), allure.Body(func() {
		}))
		allure.Step(allure.Description("Step 2"), allure.Body(func() {
			allure.Fail(errors.New("Error message"))
		}))
	}))
}

func TestFailNowAllure(t *testing.T) {
	allure.Test(t, allure.Description("Test with Allure error in it"), allure.Body(func() {
		allure.FailNow(errors.New("A more serious error"))
		allure.Step(allure.Description("Step you're not supposed to see"), allure.Body(func() {}))
	}))
}

func TestBreakAllure(t *testing.T) {
	allure.Test(t, allure.Description("Test with Allure error in it"), allure.Body(func() {
		allure.Step(allure.Description("Step 1"), allure.Body(func() {}))
		allure.Step(allure.Description("Step 2"), allure.Body(func() {
			allure.Break(errors.New("Error message"))
		}))
	}))
}

func TestBreakNowAllure(t *testing.T) {
	allure.Test(t, allure.Description("Test with Allure error in it"), allure.Body(func() {
		allure.BreakNow(errors.New("A more serious error"))
		allure.Step(allure.Description("Step you're not supposed to see"), allure.Body(func() {}))
	}))
}
