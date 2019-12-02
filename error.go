package allure

import (
	"runtime/debug"
	"strings"
	"testing"
)

func Error(err error) {
	allureError(err, false)
}

func ErrorNow(err error) {
	allureError(err, true)
}

func allureError(err error, now bool) {
	if testResult, ok := ctxMgr.GetValue(testResultKey); ok {
		testStatusDetails := testResult.(*result).StatusDetails
		if testStatusDetails == nil {
			testStatusDetails = &statusDetails{}
		}
		testStatusDetails.Trace = filterStackTrace(debug.Stack())
		testStatusDetails.Message = err.Error()
		testResult.(*result).StatusDetails = testStatusDetails
	}
	if testInstance, ok := ctxMgr.GetValue(testInstanceKey); ok {
		if now {
			testInstance.(*testing.T).FailNow()
		} else {
			testInstance.(*testing.T).Fail()
		}
	}
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
