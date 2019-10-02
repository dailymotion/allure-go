package allure_test

import (
	"testing"

	"github.com/dailymotion/allure-go"
)

func TestPassedExample(t *testing.T) {
	allure.Test(t, "This is a test to show allure implementation with a passing test", func() {
		s := "Hello world"
		if len(s) == 0 {
			t.Errorf("Expected 'hello world' string, but got %s ", s)
		}
	})
}

func TestFailedExample(t *testing.T) {
	allure.Test(t, "This is a test to show allure implementation with a failing test", func() {
		s := "Hello world"
		if len(s) != 0 {
			t.Errorf("Expected empty string, but got %s ", s)
		}
	})
}

func TestSkippedExample(t *testing.T) {
	allure.Test(t, "This is a test to show allure implemetantion with a skipped test", func() {
		t.Skip("Skipping this test")
	})
}
