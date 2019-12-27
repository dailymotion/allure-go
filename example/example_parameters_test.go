package example

import (
	"fmt"
	"github.com/dailymotion/allure-go/step"
	"github.com/dailymotion/allure-go/test"
	"testing"

	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/severity"
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
	for i := 0; i < 5; i++ {
		t.Run("", func(t *testing.T) {
			allure.Test(t,
				test.Description("Test with parameters"),
				test.Parameter("counter", i),
				test.Body(func() {}))
		})
	}
}

func TestAllureParametersExample(t *testing.T) {
	allure.Test(t,
		test.Description("This is a test to show allure implementation with a passing test"),
		test.Body(func() {
			allure.Step(step.Description(fmt.Sprintf("Number: %d", parameters["int"])), step.Action(func() {}))
			allure.Step(step.Description(fmt.Sprintf("String: %s", parameters["string"])), step.Action(func() {}))
			allure.Step(step.Description(fmt.Sprintf("Interface: %+v", parameters["struct"])), step.Action(func() {}))
		}))
}

func TestAllureStepWithParameters(t *testing.T) {
	allure.Test(t,
		test.Description("Test with steps that have parameters"),
		test.Body(func() {
			for i := 0; i < 5; i++ {
				allure.Step(
					step.Description("Step with parameters"),
					step.Parameter("counter", i),
					step.Action(func() {}))
			}
			allure.SkipStep(step.Description("Step with parameters"), step.Reason("Skip this step with parameters"), step.Parameter("counter", 6), step.Action(func() {}))
		}))
}

func TestAllureParameterTypes(t *testing.T) {
	allure.Test(t,
		test.Description("Test parameter types"),
		test.Parameter("uintptr", uintptr(10)),
		test.Parameter("uint", uint(10)),
		test.Parameter("uint8", uint8(10)),
		test.Parameter("uint16", uint16(10)),
		test.Parameter("uint32", uint32(10)),
		test.Parameter("uint64", uint64(10)),
		test.Parameter("int", int(10)),
		test.Parameter("int8", int8(10)),
		test.Parameter("int16", int16(10)),
		test.Parameter("int32", int32(10)),
		test.Parameter("int64", int64(10)),
		test.Parameter("float32", float32(10.5)),
		test.Parameter("float64", float64(10.5)),
		test.Parameter("complex64", complex(float32(2), float32(-2))),
		test.Parameter("complex128", complex(float64(2), float64(-2))),
		test.Body(func() {}))
}

func TestAllureWithLabels(t *testing.T) {
	allure.Test(t, test.Description("Test with labels"),
		test.Epic("Epic Epic of Epicness"),
		test.Lead("Duke Nukem"),
		test.Owner("Rex Powercolt"),
		test.Severity(severity.Critical),
		test.Story("story1"),
		test.Story("story2"),
		test.Feature("feature1"),
		test.Feature("feature2"),
		test.Body(func() {}))
}
