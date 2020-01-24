package example

import (
	"errors"
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestAllureSetupTeardown(t *testing.T) {
	allure.BeforeTest(t,
		allure.Description("setup"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Setup step 1"),
				allure.Action(func() {}))
		}))

	allure.Test(t,
		allure.Description("actual test"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Test step 1"),
				allure.Action(func() {}))
		}))

	allure.AfterTest(t,
		allure.Description("teardown"),
		allure.Action(func() {
			allure.Step(
				allure.Description("teardown step 1"),
				allure.Action(func() {}))
		}))
}

func TestAllureSetupFailed(t *testing.T) {
	allure.BeforeTest(t,
		allure.Description("setup failed"),
		allure.Action(func() {
			allure.Fail(errors.New("some error"))
		}))

	allure.Test(t,
		allure.Description("actual test"),
		allure.Action(func() {}))
}
