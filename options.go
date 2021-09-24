package allure

import (
	"github.com/dailymotion/allure-go/severity"
)

type Option func(r hasOptions)

func Lead(lead string) Option {
	return func(r hasOptions) {
		r.addLabel("lead", lead)
	}
}

func Owner(owner string) Option {
	return func(r hasOptions) {
		r.addLabel("owner", owner)
	}
}

func Epic(epic string) Option {
	return func(r hasOptions) {
		r.addLabel("epic", epic)
	}
}

func Severity(severity severity.Severity) Option {
	return func(r hasOptions) {
		r.addLabel("severity", string(severity))
	}
}

func Story(story string) Option {
	return func(r hasOptions) {
		r.addLabel("story", story)
	}
}

func Feature(feature string) Option {
	return func(r hasOptions) {
		r.addLabel("feature", feature)
	}
}

func Parameter(name string, value interface{}) Option {
	return func(r hasOptions) {
		r.addParameter(name, value)
	}
}

func Parameters(parameters map[string]interface{}) Option {
	return func(r hasOptions) {
		r.addParameters(parameters)
	}
}

func Description(description string) Option {
	return func(r hasOptions) {
		r.addDescription(description)
	}
}

func Action(action func()) Option {
	return func(r hasOptions) {
		r.addAction(action)
	}
}

func Reason(reason string) Option {
	return func(r hasOptions) {
		r.addReason(reason)
	}
}

func Name(name string) Option {
	return func(r hasOptions) {
		r.addName(name)
	}
}

func Suite(suite string) Option {
	return func(r hasOptions) {
		r.addLabel("suite", suite)
	}
}

func SubSuite(subSuite string) Option {
	return func(r hasOptions) {
		r.addLabel("subSuite", subSuite)
	}
}

func ParentSuite(parentSuite string) Option {
	return func(r hasOptions) {
		r.addLabel("parentSuite", parentSuite)
	}
}
