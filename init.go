package allure

import "github.com/jtolds/gls"

var (
	ctxMgr *gls.ContextManager
)

func init() {
	ctxMgr = gls.NewContextManager()
}
