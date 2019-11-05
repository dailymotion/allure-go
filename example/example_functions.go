package example

import "github.com/dailymotion/allure-go"

func doSomething() {
	allure.Step("Something", func() {
		doSomethingNested()
	})
}

func doSomethingNested() {
	allure.Step("Because we can!", func() {})
}

func addSomeParameters(parameters map[string]interface{}) {
	allure.StepWithParameter("Step with parameters", parameters, func() {})
}
