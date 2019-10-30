package example

import "github.com/dailymotion/allure-go"

func DoSomething() {
	allure.Step("Something", func() {
		DoSomethingNested()
	})
}

func DoSomethingNested() {
	allure.Step("Because we can!", func() {})
}
