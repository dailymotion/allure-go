package allure

import (
	"github.com/jtolds/gls"
	"sync"
)

var (
	ctxMgr           *gls.ContextManager
	wsd              string
	resultsPath      string
	createFolderOnce sync.Once
	copyEnvFileOnce  sync.Once
	testPhaseObjects map[string]*testPhaseContainer
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
	testPhaseObjects = make(map[string]*testPhaseContainer)
}
