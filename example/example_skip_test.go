package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestSkip(t *testing.T) {
	allure.Test(t,
		allure.Description("Skip test"),
		allure.Body(func() {
			allure.SkipStep(allure.Description("Skip"), allure.Body(func() {}), allure.Reason("reason"))
		}))
}
