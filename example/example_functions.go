package example

import (
	"github.com/dailymotion/allure-go"
)

func doSomething() {
	allure.Step(allure.Description("Something"), allure.Body(func() {
		doSomethingNested()
	}))
}

func doSomethingNested() {
	allure.Step(allure.Description("Because we can!"), allure.Body(func() {}))
}

func addSomeParameters(parameters map[string]interface{}) {
	var stepOptions = make([]allure.Option, 0)
	for k, v := range parameters {
		stepOptions = append(stepOptions, allure.Parameter(k, v))
	}
	stepOptions = append(stepOptions, allure.Description("Step with parameters"))
	stepOptions = append(stepOptions, allure.Body(func() {}))

	allure.Step(stepOptions...)
}
