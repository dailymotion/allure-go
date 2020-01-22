package allure

import (
	"github.com/jtolds/gls"
	"sync"
)

var (
	ctxMgr           *gls.ContextManager
	wsd              string
	ResultsPath      string
	CreateFolderOnce sync.Once
	CopyEnvFileOnce  sync.Once
	testPhaseObjects map[string]*testPhaseContainer
)

const (
	ResultsPathEnvKey = "ALLURE_RESULTS_PATH"
	WsPathEnvKey      = "ALLURE_WORKSPACE_PATH"
	EnvFileKey        = "ALLURE_ENVIRONMENT_FILE_PATH"
	nodeKey           = "current_step_container"
	testResultKey     = "test_result_object"
	testInstanceKey   = "test_instance"
)

func init() {
	ctxMgr = gls.NewContextManager()
	testPhaseObjects = make(map[string]*testPhaseContainer)
}
