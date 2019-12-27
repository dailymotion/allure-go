package options

type HasOptions interface {
	AddLabel(key string, value string)
	AddDescription(description string)
	AddParameter(name string, value interface{})
	AddName(name string)
	AddAction(action func())
	AddReason(reason string)
}
