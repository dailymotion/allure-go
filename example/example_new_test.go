package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestNewTest(t *testing.T) {
	allure.Test(
		t,
		allure.Description("New Test Description"),
		allure.Body(func() {
			allure.Step(
				allure.Description("Step description"),
				allure.Body(func() {

				}))
		}))
}
