package main

import (
	"testing"
)

func TestSample(t *testing.T) {
    testSample := 1

    if testSample == 1 {
        t.Log("success")
    } else {
        t.Error("fail")
    }
}
