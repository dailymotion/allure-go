package test

import (
	"github.com/dailymotion/allure-go/options"
	"github.com/dailymotion/allure-go/severity"
)

type Option func(r *options.HasOptions)

func Lead(lead string) Option {
	return func(r *options.HasOptions) {
		r.AddLabel("lead", lead)
	}
}

func Owner(owner string) Option {
	return func(r *options.HasOptions) {
		r.AddLabel("owner", owner)
	}
}

func Epic(epic string) Option {
	return func(r *options.HasOptions) {
		r.AddLabel("epic", epic)
	}
}

func Severity(severity severity.Severity) Option {
	return func(r *options.HasOptions) {
		r.AddLabel("severity", string(severity))
	}
}

func Story(story string) Option {
	return func(r *options.HasOptions) {
		r.AddLabel("story", story)
	}
}

func Feature(feature string) Option {
	return func(r *options.HasOptions) {
		r.AddLabel("feature", feature)
	}
}

func TestParameter(name string, value interface{}) Option {
	return func(r *options.HasOptions) {
		r.AddParameter(name, value)
	}
}

func Description(description string) Option {
	return func(r *options.HasOptions) {
		r.AddDescription(description)
	}
}

func Body(action func()) Option {
	return func(r *options.HasOptions) {
		r.AddAction(action)
	}
}
