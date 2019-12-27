package allure

type hasAttachments interface {
	GetAttachments() []Attachment
	AddAttachment(attachment Attachment)
}
