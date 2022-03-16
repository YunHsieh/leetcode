package main

import (
	"testing"
)


func TestValidateStackSequencesSence1(t *testing.T) {

	n := ValidateStackSequences([]int{1,2,3,4,5}, []int{4,5,3,2,1})
	if n {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
