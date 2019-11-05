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
