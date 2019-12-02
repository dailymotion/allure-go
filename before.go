package allure

type Container struct {
	UUID        string            `json:"uuid"`
	Name        string            `json:"name"`
	Children    []string          `json:"children"`
	Description string            `json:"description"`
	Befores     []helperContainer `json:"befores"`
	Afters      []helperContainer `json:"afters"`
	Links       []string          `json:"links"`
	Start       int64             `json:"start"`
	Stop        int64             `json:"stop"`
}

func (container Container) writeResultsFile() error {
	//TODO: implement that
	return nil
}

//helperContainer defines a step
type helperContainer struct {
	Name          string         `json:"name,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetails *statusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage,omitempty"`
	Description   string         `json:"description,omitempty"`
	Start         int64          `json:"start,omitempty"`
	Stop          int64          `json:"stop,omitempty"`
	Steps         []stepObject   `json:"steps,omitempty"`
	Attachments   []attachment   `json:"attachments,omitempty"`
	Parameters    []Parameter    `json:"parameters,omitempty"`
}

func newHelper() *helperContainer {
	return &helperContainer{
		Steps:       make([]stepObject, 0),
		Attachments: make([]attachment, 0),
		Parameters:  make([]Parameter, 0),
	}
}
