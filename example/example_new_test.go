package example

import (
	"fmt"
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/severity"
	"testing"
)

func TestNewTest(t *testing.T) {
	allure.NewTest(t,
		allure.Description("Test"),
		allure.Severity(severity.Normal),
		allure.Body(func() {
			allure.NewStep("Stuff", allure.Action(func() {
				fmt.Println("dfhdkjhdfkjlsf")
			}))
		}))
}
