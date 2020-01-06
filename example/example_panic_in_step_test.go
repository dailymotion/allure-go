package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestPanicInStep(t *testing.T) {
	allure.Test(t,
		allure.Description("Panic handling test"),
		allure.Body(func() {
			allure.Step(
				allure.Description("step that will panic"),
				allure.Body(func() {
					panic("throwing a panic")
				}))
		}))
}
