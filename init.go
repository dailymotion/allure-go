package allure

import (
	"sync"

	"github.com/jtolds/gls"
)

var (
	ctxMgr           *gls.ContextManager
	wsd              string
	resultsPath      string
	createFolderOnce sync.Once
	copyEnvFileOnce  sync.Once
	testPhaseObjects sync.Map
)

const (
	resultsPathEnvKey = "ALLURE_RESULTS_PATH"
	wsPathEnvKey      = "ALLURE_WORKSPACE_PATH"
	envFileKey        = "ALLURE_ENVIRONMENT_FILE_PATH"
	nodeKey           = "current_step_container"
	testResultKey     = "test_result_object"
	testInstanceKey   = "test_instance"
)

func init() {
	ctxMgr = gls.NewContextManager()
	testPhaseObjects = sync.Map{}
}
