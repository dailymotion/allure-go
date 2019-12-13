package allure

type hasStatusDeatils interface {
	getStatusDetails() *statusDetails
	setStatusDetails(details statusDetails)
}
