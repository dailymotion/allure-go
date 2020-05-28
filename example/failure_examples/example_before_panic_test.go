package failure_examples

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestBeforePanic(t *testing.T) {
	allure.BeforeTest(t, allure.Action(func() {
		panic("panic at the before statement! (disco)")
	}))

	allure.Test(t, allure.Action(func() {}))
}
