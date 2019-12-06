package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestAllureSetupTeardown(t *testing.T) {
	allure.BeforeTest(t, "setup", func() {
		allure.Step("Setup step 1", func() {})
	})

	allure.Test(t, "actual test", func() {
		allure.Step("Test step 1", func() {})
	})

	allure.AfterTest(t, "teardown", func() {
		allure.Step("teardown step 1", func() {})
	})
}
