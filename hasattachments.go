package allure

type hasAttachments interface {
	getAttachments() []attachment
	addAttachment(attachment attachment)
}
