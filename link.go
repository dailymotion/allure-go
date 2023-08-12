package allure

type LinkType string

const (
	IssueType    LinkType = "issue"
	AnyLinkType  LinkType = "link"
	TestCaseType LinkType = "test_case"
)

type link struct {
	Url  string   `json:"url,omitempty"`
	Name string   `json:"name,omitempty"`
	Type LinkType `json:"type,omitempty"`
}
