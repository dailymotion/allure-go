package failure_examples

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestPanicBasic(t *testing.T) {
	allure.Test(t, allure.Description("Panic handling test"),
		allure.Action(func() {
			panic("throwing a panic")
		}))
}
