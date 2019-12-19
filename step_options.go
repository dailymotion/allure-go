package allure

type StepOption func(s *stepObject)

func AddStepParameter(name string, value interface{}) StepOption {
	return func(s *stepObject) {
		parameter := parseParameter(name, value)
		s.Parameters = append(s.Parameters, parameter)
	}
}

func Action(action func()) StepOption {
	return func(s *stepObject) {
		s.action = action
	}
}
