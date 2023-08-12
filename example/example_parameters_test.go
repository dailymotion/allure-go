package example

import (
	"fmt"
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
				allure.Description("Test with parameters"),
				allure.Parameter("counter", i),
				allure.Action(func() {}))
		})
	}
}

func TestAllureParametersExample(t *testing.T) {
	allure.Test(t,
		allure.Parameters(parameters),
		allure.Description("This is a test to show allure implementation with a passing test"),
		allure.Action(func() {
			allure.Step(allure.Description(fmt.Sprintf("Number: %d", parameters["int"])), allure.Action(func() {}))
			allure.Step(allure.Description(fmt.Sprintf("String: %s", parameters["string"])), allure.Action(func() {}))
			allure.Step(allure.Description(fmt.Sprintf("Interface: %+v", parameters["struct"])), allure.Action(func() {}))
		}))
}

func TestAllureStepWithParameters(t *testing.T) {
	allure.Test(t,
		allure.Description("Test with steps that have parameters"),
		allure.Action(func() {
			for i := 0; i < 5; i++ {
				allure.Step(
					allure.Description("Step with parameters"),
					allure.Parameter("counter", i),
					allure.Action(func() {}))
			}
			allure.SkipStep(allure.Description("Step with parameters"), allure.Reason("Skip this step with parameters"), allure.Parameter("counter", 5), allure.Action(func() {}))
		}))
}

func TestAllureParameterTypes(t *testing.T) {
	allure.Test(t,
		allure.Description("Test parameter types"),
		allure.Parameter("uintptr", uintptr(10)),
		allure.Parameter("uint", uint(10)),
		allure.Parameter("uint8", uint8(10)),
		allure.Parameter("uint16", uint16(10)),
		allure.Parameter("uint32", uint32(10)),
		allure.Parameter("uint64", uint64(10)),
		allure.Parameter("int", int(10)),
		allure.Parameter("int8", int8(10)),
		allure.Parameter("int16", int16(10)),
		allure.Parameter("int32", int32(10)),
		allure.Parameter("int64", int64(10)),
		allure.Parameter("float32", float32(10.5)),
		allure.Parameter("float64", float64(10.5)),
		allure.Parameter("complex64", complex(float32(2), float32(-2))),
		allure.Parameter("complex128", complex(float64(2), float64(-2))),
		allure.Action(func() {}))
}

func TestAllureWithLabels(t *testing.T) {
	allure.Test(t, allure.Description("Test with labels"),
		allure.Epic("Epic Epic of Epicness"),
		allure.ID("id1"),
		allure.Lead("Duke Nukem"),
		allure.Owner("Rex Powercolt"),
		allure.Severity(severity.Critical),
		allure.Story("story1"),
		allure.Story("story2"),
		allure.Feature("feature1"),
		allure.Feature("feature2"),
		allure.Layer("integration-tests"),
		allure.Tag("tag1"),
		allure.Tags("tag2", "tag3"),
		allure.Label("customLabel1", "customLabel1Value"),
		allure.Action(func() {}))
}

func TestAllureNamedTest(t *testing.T) {
	allure.Test(t,
		allure.Description("Description of a test"),
		allure.Name("Test naming of a test"),
		allure.Action(func() {
			allure.Step()
		}))
}
