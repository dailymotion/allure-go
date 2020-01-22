package allure

type hasOptions interface {
	addLabel(key string, value string)
	addDescription(description string)
	addParameter(name string, value interface{})
	addParameters(parameters map[string]interface{})
	addName(name string)
	addAction(action func())
	addReason(reason string)
}

// This interface provides functions required to manipulate children step records, used in the result object and
// step object for recursive handling
type hasSteps interface {
	getSteps() []stepObject
	addStep(step stepObject)
}

type hasStatus interface {
	setStatus(status string)
	getStatus() string
}

type hasAttachments interface {
	getAttachments() []attachment
	addAttachment(attachment attachment)
}

type hasStatusDetails interface {
	getStatusDetails() *statusDetails
	setStatusDetails(details statusDetails)
}
