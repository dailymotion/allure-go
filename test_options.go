package allure

import (
	"github.com/dailymotion/allure-go/severity"
)

type Option func(r *result)

func Lead(lead string) Option {
	return func(r *result) {
		r.addLabel("lead", lead)
	}
}

func Owner(owner string) Option {
	return func(r *result) {
		r.addLabel("owner", owner)
	}
}

func Epic(epic string) Option {
	return func(r *result) {
		r.addLabel("epic", epic)
	}
}

func Severity(severity severity.Severity) Option {
	return func(r *result) {
		r.addLabel("severity", string(severity))
	}
}

func Story(story string) Option {
	return func(r *result) {
		r.addLabel("story", story)
	}
}

func Feature(feature string) Option {
	return func(r *result) {
		r.addLabel("feature", feature)
	}
}

func AddParameter(name string, value interface{}) Option {
	return func(r *result) {
		parameter := parseParameter(name, value)
		r.Parameters = append(r.Parameters, parameter)
	}
}

func Description(description string) Option {
	return func(r *result) {
		r.Description = description
	}
}

func Body(action func()) Option {
	return func(r *result) {
		r.test = action
	}
}
