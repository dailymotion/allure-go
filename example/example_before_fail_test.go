package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestBeforeFail(t *testing.T) {
	allure.BeforeTest(t, allure.Action(func() {
		panic("panic at the before statement! (disco)")
	}))

	allure.Test(t, allure.Action(func() {}))
}
