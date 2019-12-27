package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/step"
)

func doSomething() {
	allure.Step(step.Description("Something"), step.Action(func() {
		doSomethingNested()
	}))
}

func doSomethingNested() {
	allure.Step(step.Description("Because we can!"), step.Action(func() {}))
}

func addSomeParameters(parameters map[string]interface{}) {
	var stepOptions = make([]step.Option, 0)
	for k, v := range parameters {
		stepOptions = append(stepOptions, step.Parameter(k, v))
	}
	stepOptions = append(stepOptions, step.Description("Step with parameters"))
	stepOptions = append(stepOptions, step.Action(func() {}))

	allure.Step(stepOptions...)
}
