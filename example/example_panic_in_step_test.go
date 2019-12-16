package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestPanicInStep(t *testing.T) {
	allure.Test(t, "Panic handling test", func() {
		allure.Step("step that will panic", func() {
			panic("throwing a panic")
		})
	})
}
