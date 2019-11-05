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
)

func init() {
	ctxMgr = gls.NewContextManager()
}
