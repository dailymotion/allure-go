package example

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestAllureWithLinks(t *testing.T) {
	allure.Test(t, allure.Description("Test with links"),
		allure.Link("https://github.com/", "GitHub Link"),
		allure.Issue("https://github.com/", "GitHub Issue"),
		allure.TestCase("https://github.com/", "GitHub TestCase"),
		allure.Action(func() {}))
}
