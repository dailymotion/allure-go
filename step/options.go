package step

import "github.com/dailymotion/allure-go"

type Option func(s *allure.StepObject)

func Parameter(name string, value interface{}) Option {
	return func(s *allure.StepObject) {
		parameter := allure.ParseParameter(name, value)
		s.Parameters = append(s.Parameters, parameter)
	}
}

func Action(action func()) Option {
	return func(s *allure.StepObject) {
		s.Action = action
	}
}

func Description(description string) Option {
	return func(s *allure.StepObject) {
		s.Name = description
	}
}
