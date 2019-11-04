package allure

import (
	"fmt"
	"github.com/jtolds/gls"
	"log"
	"os"
)

var (
	ctxMgr     *gls.ContextManager
	wsd        string
	resultPath string
)

const (
	resultsPathEnvKey = "ALLURE_RESULTS_PATH"
	nodeKey           = "current_step_container"
)

func init() {
	ctxMgr = gls.NewContextManager()
	wsd = os.Getenv(resultsPathEnvKey)
	if wsd == "" {
		log.Fatalf("%s environment variable cannot be empty", resultsPathEnvKey)
	}
	resultPath = fmt.Sprintf("%s/allure-results", wsd)
}
