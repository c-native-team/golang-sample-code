package main

import (
	"testing"
)

func Test_input_string(t *testing.T){
	output := input_string("Test")
	if output != "Test"{
		t.Errorf("Calculate(\"Test\") = %v, expected Test", output)
	}
}
