package hello

import "testing"

func Test_Hello(t *testing.T) {
	if Hello() != "hello world" {
		t.Error("wrong result")
	}
}
