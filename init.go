package allure

import (
	"github.com/jtolds/gls"
)

var (
	ctxMgr     *gls.ContextManager
	wsd        string
	resultPath string
)

const (
	resultsPathEnvKey = "ALLURE_RESULTS_PATH"
	wsPathEnvKey      = "ALLURE_WORKSPACE_PATH"
	nodeKey           = "current_step_container"
	testResultKey     = "test_result_object"
	testInstanceKey   = "test_instance"
)

func init() {
	ctxMgr = gls.NewContextManager()
}
