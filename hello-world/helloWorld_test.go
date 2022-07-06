package helloWorld

import "testing"

func TestGetHelloWorld(t *testing.T) {
	if GetHelloWorld() != Text {
		t.Error("Wrong text")
	}
}
