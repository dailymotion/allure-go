package test

import (
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/severity"
)

type Option func(r *allure.Result)

func Lead(lead string) Option {
	return func(r *allure.Result) {
		r.AddLabel("lead", lead)
	}
}

func Owner(owner string) Option {
	return func(r *allure.Result) {
		r.AddLabel("owner", owner)
	}
}

func Epic(epic string) Option {
	return func(r *allure.Result) {
		r.AddLabel("epic", epic)
	}
}

func Severity(severity severity.Severity) Option {
	return func(r *allure.Result) {
		r.AddLabel("severity", string(severity))
	}
}

func Story(story string) Option {
	return func(r *allure.Result) {
		r.AddLabel("story", story)
	}
}

func Feature(feature string) Option {
	return func(r *allure.Result) {
		r.AddLabel("feature", feature)
	}
}

func Parameter(name string, value interface{}) Option {
	return func(r *allure.Result) {
		parameter := allure.ParseParameter(name, value)
		r.Parameters = append(r.Parameters, parameter)
	}
}

func Description(description string) Option {
	return func(r *allure.Result) {
		r.Description = description
	}
}

func Body(action func()) Option {
	return func(r *allure.Result) {
		r.Test = action
	}
}
