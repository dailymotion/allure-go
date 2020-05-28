package allure

import (
	"runtime/debug"
	"strings"
	"testing"
)

// Fail sets the current step as well as the entire test script as failed and stores the stack trace in the result object.
// Script execution is not interrupted, script would still run after this call.
func Fail(err error) {
	allureError(err, "failed", false)
}

// FailNow sets the current step as well as the entire test script as failed and stores the stack trace in the result object.
// Script execution is interrupted, script stops immediately.
func FailNow(err error) {
	allureError(err, "failed", true)
}

// Break sets the current step as well as the entire test script as broken and stores the stack trace in the result object.
// Script execution is not interrupted, script would still run after this call.
func Break(err error) {
	allureError(err, "broken", false)
}

// BreakNow sets the current step as well as the entire test script as broken and stores the stack trace in the result object.
// Script execution is interrupted, script stops immediately.
func BreakNow(err error) {
	allureError(err, "broken", true)
}

func allureError(err error, status string, now bool) {
	manipulateOnObjectFromCtx(
		testResultKey,
		func(testResult interface{}) {
			testStatusDetails := testResult.(hasStatusDetails).getStatusDetails()
			if testStatusDetails == nil {
				testStatusDetails = &statusDetails{}
			}
			testStatusDetails.Trace = filterStackTrace(debug.Stack())
			testStatusDetails.Message = err.Error()
			testResult.(hasStatusDetails).setStatusDetails(*testStatusDetails)
			testResult.(hasStatus).setStatus(status)
		})
	manipulateOnObjectFromCtx(
		nodeKey,
		func(node interface{}) {
			node.(hasStatus).setStatus(status)
		})
	manipulateOnObjectFromCtx(
		testInstanceKey,
		func(testInstance interface{}) {
			testInstance.(*testing.T).Fail()
			if now {
				panic(err)
			}
		})
}

func filterStackTrace(stack []byte) string {
	stringTraces := strings.Split(string(stack), "\n")
	result := stringTraces[0] + "\n"
	for i := 1; i+1 < len(stringTraces); i = i + 2 {
		// for vendored code calls
		if !strings.Contains(stringTraces[i+1], "allure-go/vendor/") &&
			// for allure-go specific function calls
			!strings.HasPrefix(stringTraces[i], "github.com/dailymotion/allure-go.") {
			result += stringTraces[i] + "\n" + stringTraces[i+1] + "\n"
		}
	}

	return result
}
