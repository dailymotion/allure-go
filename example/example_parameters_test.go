package example_test

import (
	"fmt"
	"github.com/dailymotion/allure-go"
	"testing"
)

type SampleObject struct {
	Key   string
	Value string
}

var sampleStruct = SampleObject{
	Key:   "some key",
	Value: "some value",
}

var parameters = map[string]interface{}{
	"string": "value",
	"int":    10,
	"struct": sampleStruct,
}

func TestAllureParameterized(t *testing.T) {
	allure.TestWithParameters(t,
		"This is a test to show allure implementation with a passing test",
		parameters,
		func() {
			allure.Step(fmt.Sprintf("Number: %d", parameters["int"]), func() {})
			allure.Step(fmt.Sprintf("String: %s", parameters["string"]), func() {})
			allure.Step(fmt.Sprintf("Interface: %+v", parameters["struct"]), func() {})
		})
}

func TestAllureParameterizedFailing(t *testing.T) {
	allure.TestWithParameters(t,
		"This is a test to show allure implementation with a passing test",
		parameters,
		func() {
			allure.Step(fmt.Sprintf("Int: %d; String: %s; Interface: %+v", parameters["int"], parameters["string"], parameters["struct"]), func() {})
			s := "Hello world"
			if len(s) == 0 {
				t.Errorf("Expected 'hello world' string, but got %s ", s)
			}
		})
}
