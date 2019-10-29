package allure

type context struct {
}

func getCurrentContext() *context {
	ctx, ok := ctxMgr.GetValue(nodeKey)
	if ok {
		return ctx.(*context)
	}
	return nil
}

type A interface {
	Step(items ...interface{})
	Attach(items ...interface{})
}

func (ctx *context) Step(items ...interface{}) {

}

func (ctx *context) Attach(items ...interface{}) {

}

func (ctx *context) allureStepInner(f func(A)) {
	f(ctx)
}
