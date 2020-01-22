package example

import (
	"fmt"
	"testing"
)

func TestPanicking(t *testing.T) {
	panic("")
}

func TestOkay(t *testing.T) {
	fmt.Println("okay")
}
