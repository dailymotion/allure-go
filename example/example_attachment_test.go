package example

import (
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/step"
	"github.com/dailymotion/allure-go/test"
	"io/ioutil"
	"log"
	"testing"
)

func TestTextAttachmentToStep(t *testing.T) {
	allure.Test(t, test.Description("Testing a text attachment"), test.Body(func() {
		allure.Step(step.Description("adding a text attachment"), step.Action(func() {
			_ = allure.AddAttachment("text!", allure.TextPlain, []byte("Some text!"))
		}))
	}))
}

func TestImageAttachmentToStep(t *testing.T) {
	allure.Test(t, test.Description("testing an image attachment"), test.Body(func() {
		allure.Step(step.Description("adding an image attachment"), step.Action(func() {
			dat, err := ioutil.ReadFile("../Coryphaena_hippurus.png")
			if err != nil {
				log.Println(err)
			}
			_ = allure.AddAttachment("mahi-mahi", allure.ImagePng, dat)
		}))
	}))
}

func TestTextAttachment(t *testing.T) {
	allure.Test(t, test.Description("Testing a text attachment"), test.Body(func() {
		_ = allure.AddAttachment("text!", allure.TextPlain, []byte("Some text!"))
	}))
}

func TestImageAttachment(t *testing.T) {
	allure.Test(t, test.Description("testing an image attachment"), test.Body(func() {
		dat, err := ioutil.ReadFile("../Coryphaena_hippurus.png")
		if err != nil {
			log.Println(err)
		}
		_ = allure.AddAttachment("mahi-mahi", allure.ImagePng, dat)
	}))
}
