package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestAllureSetupTeardown(t *testing.T) {
	allure.Before(t, "setup", func() {
		allure.Step("Setup step 1", func() {})
	})

	allure.Test(t, "actual test", func() {
		allure.Step("Test step 1", func() {})
	})

	allure.After(t, "teardown", func() {
		allure.Step("teardown step 1", func() {})
	})
}
