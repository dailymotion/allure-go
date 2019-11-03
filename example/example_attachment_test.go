package example_test

import (
	"github.com/dailymotion/allure-go"
	"io/ioutil"
	"log"
	"testing"
)

func TestTextAttachmentToStep(t *testing.T) {
	allure.Test(t, "Testing a text attachment", func() {
		allure.Step("adding a text attachment", func() {
			_ = allure.AddTextAttachment("text!", allure.TextPlain, "Some text!")
		})
	})
}

func TestImageAttachmentToStep(t *testing.T) {
	allure.Test(t, "testing an image attachment", func() {
		allure.Step("adding an image attachment", func() {
			dat, err := ioutil.ReadFile("../Coryphaena_hippurus.png")
			if err != nil {
				log.Fatal(err)
			}
			_ = allure.AddByteArrayAttachment("mahi-mahi", allure.ImagePng, dat)
		})
	})
}

func TestTextAttachment(t *testing.T) {
	allure.Test(t, "Testing a text attachment", func() {
		_ = allure.AddTextAttachment("text!", allure.TextPlain, "Some text!")
	})
}

func TestImageAttachment(t *testing.T) {
	allure.Test(t, "testing an image attachment", func() {
		dat, err := ioutil.ReadFile("../Coryphaena_hippurus.png")
		if err != nil {
			log.Fatal(err)
		}
		_ = allure.AddByteArrayAttachment("mahi-mahi", allure.ImagePng, dat)
	})
}
