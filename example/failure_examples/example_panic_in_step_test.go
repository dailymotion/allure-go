package failure_examples

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestPanicInStep(t *testing.T) {
	allure.Test(t,
		allure.Description("Panic handling test"),
		allure.Action(func() {
			allure.Step(
				allure.Description("step that will panic"),
				allure.Action(func() {
					panic("throwing a panic")
				}))
		}))
}
