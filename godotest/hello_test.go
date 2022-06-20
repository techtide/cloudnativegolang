package main

import "testing"

func TestHello(t *testing.T) {
	emptyResult := hello("")

	if emptyResult != "Hello unknown user!" {
		t.Error("it failed :(")
	}
}
