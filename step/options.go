package step

import (
	"github.com/dailymotion/allure-go/options"
)

type Option func(s *options.HasOptions)

func Parameter(name string, value interface{}) Option {
	return func(s *options.HasOptions) {
		s.AddParameter(name, value)
	}
}

func Action(action func()) Option {
	return func(s *options.HasOptions) {
		s.AddAction(action)
	}
}

func Description(description string) Option {
	return func(s *options.HasOptions) {
		s.AddName(description)
	}
}
