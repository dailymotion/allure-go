package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestStuff(t *testing.T) {
	allure.Test(t,
		allure.Description("Verifying copy/paste feature works"),
		allure.Action(func() {
			allure.Step(allure.Description("Container step"),
				allure.Action(func() {
					step := allure.CopyStep(allure.Description("Copied step"))
					for i := 0; i < 10; i++ {
						allure.PasteStep(step)
					}
				}))
		}))
}
