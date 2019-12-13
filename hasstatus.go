package allure

type hasStatus interface {
	setStatus(status string)
	getStatus() string
}
