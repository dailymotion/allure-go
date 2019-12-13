package allure

// This interface provides functions required to manipulate children step records, used in the result object and
// step object for recursive handling
type hasSteps interface {
	getSteps() []stepObject
	addStep(step stepObject)
}
