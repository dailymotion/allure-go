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
			allure.TestWithParameters(t,
				"Test with parameters",
				map[string]interface{}{
					"counter": i,
				},
				nil,
				func() {

				})
		})
	}
}

func TestAllureParametersExample(t *testing.T) {
	allure.TestWithParameters(t,
		"This is a test to show allure implementation with a passing test",
		parameters,
		nil,
		func() {
			allure.Step(fmt.Sprintf("Number: %d", parameters["int"]), func() {})
			allure.Step(fmt.Sprintf("String: %s", parameters["string"]), func() {})
			allure.Step(fmt.Sprintf("Interface: %+v", parameters["struct"]), func() {})
		})
}

func TestAllureStepWithParameters(t *testing.T) {
	allure.Test(t, "Test with steps that have parameters", func() {
		for i := 0; i < 5; i++ {
			allure.StepWithParameter("Step with parameters", map[string]interface{}{"counter": i}, func() {})
		}
		allure.SkipStepWithParameter("Step with parameters", "Skip this step with parameters", map[string]interface{}{"counter": 6}, func() {})
	})
}

func TestAllureParameterTypes(t *testing.T) {
	allure.TestWithParameters(t,
		"Test parameter types",
		map[string]interface{}{
			"uintptr":    uintptr(10),
			"uint":       uint(10),
			"uint8":      uint8(10),
			"uint16":     uint16(10),
			"uint32":     uint32(10),
			"uint64":     uint64(10),
			"int":        int(10),
			"int8":       int8(10),
			"int16":      int16(10),
			"int32":      int32(10),
			"int64":      int64(10),
			"float32":    float32(10.5),
			"float64":    float64(10.5),
			"complex64":  complex(float32(2), float32(-2)),
			"complex128": complex(float64(2), float64(-2)),
		},
		nil,
		func() {})
}

func TestAllureWithLabels(t *testing.T) {
	allure.TestWithParameters(t, "Test with labels",
		nil,
		&allure.TestLabels{
			Epic:     "Epic Epic of Epicness",
			Lead:     "Duke Nukem",
			Owner:    "Rex Powercolt",
			Severity: severity.Critical,
			Story:    []string{"story1", "story2"},
			Feature:  []string{"feature1", "feature2"},
		}, func() {

		})
}
