package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/pkg/errors"
	"testing"
)

func TestErrorAllure(t *testing.T) {
	allure.Test(t, "Test with Allure error in it", func() {
		allure.Step("Step 1", func() {
		})
		allure.Step("Step 2", func() {
			allure.Error(errors.New("Error message"))
		})
	})
}

func TestErrorNowAllure(t *testing.T) {
	allure.Test(t, "Test with Allure error in it", func() {
		allure.ErrorNow(errors.New("A more serious error"))
		allure.Step("Step you're not supposed to see", func() {})
	})
}
