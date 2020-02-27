package allure

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
)

type attachment struct {
	uuid    string
	Name    string   `json:"name"`
	Source  string   `json:"source"`
	Type    MimeType `json:"type"`
	content []byte
}

type MimeType string

func AddAttachment(name string, mimeType MimeType, content []byte) error {
	attachment := newAttachment(name, mimeType, content)
	err := attachment.writeAttachmentFile()
	if err != nil {
		return errors.Wrap(err, "Failed to create an attachment file")
	}
	if node, ok := ctxMgr.GetValue(nodeKey); ok {
		node.(hasAttachments).addAttachment(*attachment)
	}

	return nil
}

func (a *attachment) writeAttachmentFile() error {
	resultsPathEnv := os.Getenv(resultsPathEnvKey)
	ensureFolderCreated()
	if resultsPathEnv == "" {
		log.Printf("%s environment variable cannot be empty\n", resultsPathEnvKey)
	}
	if resultsPath == "" {
		resultsPath = fmt.Sprintf("%s/allure-results", resultsPathEnv)
	}

	a.Source = fmt.Sprintf("%s-attachment", a.uuid)
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s-attachment", resultsPath, a.uuid), a.content, 0777)
	if err != nil {
		return errors.Wrap(err, "Failed to write in file")
	}
	return nil
}

func newAttachment(name string, mimeType MimeType, content []byte) *attachment {
	result := &attachment{
		uuid:    generateUUID(),
		content: content,
		Name:    name,
		Type:    mimeType,
	}

	return result
}
