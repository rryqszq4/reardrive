package core

import (
	"fmt"
	"testing"
)

func p() {
	fmt.Println(123)
}

func TestSignal(t *testing.T) {
	s:=&SignalT{QuitHandle:p}
	s.Create()
}
