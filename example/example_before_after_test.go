package example

import (
	"errors"
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

func TestAllureSetupFailed(t *testing.T) {
	allure.BeforeTest(t, "setup failed", func() {
		allure.Fail(errors.New("some error"))
	})

	allure.Test(t, "actual test", func() {})
}
