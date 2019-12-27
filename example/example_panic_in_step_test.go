package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/step"
	"github.com/dailymotion/allure-go/test"
	"testing"
)

func TestPanicInStep(t *testing.T) {
	allure.Test(t,
		test.Description("Panic handling test"),
		test.Body(func() {
			allure.Step(
				step.Description("step that will panic"),
				step.Action(func() {
					panic("throwing a panic")
				}))
		}))
}
