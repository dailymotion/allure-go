package allure

import (
	"github.com/jtolds/gls"
	"sync"
)

var (
	ctxMgr      *gls.ContextManager
	wsd         string
	resultsPath string
	once        sync.Once
)

const (
	resultsPathEnvKey = "ALLURE_RESULTS_PATH"
	wsPathEnvKey      = "ALLURE_WORKSPACE_PATH"
	envFileKey        = "ALLURE_ENVIRONMENT_FILE_PATH"
	nodeKey           = "current_step_container"
)

func init() {
	ctxMgr = gls.NewContextManager()
}
