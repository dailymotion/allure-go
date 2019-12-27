package allure

type hasStatus interface {
	SetStatus(status string)
	GetStatus() string
}
