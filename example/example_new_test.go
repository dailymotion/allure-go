package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/step"
	"github.com/dailymotion/allure-go/test"
	"testing"
)

func TestNewTest(t *testing.T) {
	allure.Test(
		t,
		test.Description("New Test Description"),
		test.Body(func() {
			allure.Step(
				step.Description("Step description"),
				step.Action(func() {

				}))
		}))
}
