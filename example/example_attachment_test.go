package example

import (
	"github.com/dailymotion/allure-go"
	"io/ioutil"
	"log"
	"testing"
)

func TestTextAttachmentToStep(t *testing.T) {
	allure.Test(t, allure.Description("Testing a text attachment"), allure.Action(func() {
		allure.Step(allure.Description("adding a text attachment"), allure.Action(func() {
			_ = allure.AddAttachment("text!", allure.TextPlain, []byte("Some text!"))
		}))
	}))
}

func TestImageAttachmentToStep(t *testing.T) {
	allure.Test(t, allure.Description("testing an image attachment"), allure.Action(func() {
		allure.Step(allure.Description("adding an image attachment"), allure.Action(func() {
			dat, err := ioutil.ReadFile("../Coryphaena_hippurus.png")
			if err != nil {
				log.Println(err)
			}
			_ = allure.AddAttachment("mahi-mahi", allure.ImagePng, dat)
		}))
	}))
}

func TestTextAttachment(t *testing.T) {
	allure.Test(t, allure.Description("Testing a text attachment"), allure.Action(func() {
		_ = allure.AddAttachment("text!", allure.TextPlain, []byte("Some text!"))
	}))
}

func TestImageAttachment(t *testing.T) {
	allure.Test(t, allure.Description("testing an image attachment"), allure.Action(func() {
		dat, err := ioutil.ReadFile("../Coryphaena_hippurus.png")
		if err != nil {
			log.Println(err)
		}
		_ = allure.AddAttachment("mahi-mahi", allure.ImagePng, dat)
	}))
}
