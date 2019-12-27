package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/test"
	"testing"
)

func TestPanicBasic(t *testing.T) {
	allure.Test(t, test.Description("Panic handling test"),
		test.Body(func() {
			panic("throwing a panic")
		}))
}
