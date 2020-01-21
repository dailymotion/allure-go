package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestSkip(t *testing.T) {
	allure.Test(t,
		allure.Description("Skip test"),
		allure.Action(func() {
			allure.SkipStep(allure.Description("Skip"), allure.Action(func() {}), allure.Reason("reason"))
		}))
}
