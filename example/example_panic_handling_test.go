package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestPanicBasic(t *testing.T) {
	allure.Test(t, "Panic handling test", func() {
		panic("throwing a panic")
	})
}
