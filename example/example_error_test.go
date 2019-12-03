package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/pkg/errors"
	"testing"
)

func TestFailAllure(t *testing.T) {
	allure.Test(t, "Test with Allure error in it", func() {
		allure.Step("Step 1", func() {
		})
		allure.Step("Step 2", func() {
			allure.Fail(errors.New("Error message"))
		})
	})
}

func TestFailNowAllure(t *testing.T) {
	allure.Test(t, "Test with Allure error in it", func() {
		allure.FailNow(errors.New("A more serious error"))
		allure.Step("Step you're not supposed to see", func() {})
	})
}

func TestBreakAllure(t *testing.T) {
	allure.Test(t, "Test with Allure error in it", func() {
		allure.Step("Step 1", func() {
		})
		allure.Step("Step 2", func() {
			allure.Break(errors.New("Error message"))
		})
	})
}

func TestBreakNowAllure(t *testing.T) {
	allure.Test(t, "Test with Allure error in it", func() {
		allure.BreakNow(errors.New("A more serious error"))
		allure.Step("Step you're not supposed to see", func() {})
	})
}
