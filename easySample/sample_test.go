package main

import (
	"testing"
)

func TestSampleScence1(t *testing.T) {
    result := testFunc()
    ans := 1

    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail: ", result)
    }
}

func testFunc() int {
    return 1
}

func TestSampleScence2(t *testing.T) {
    result := fizzBuzz(15)
    ans := "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz"
    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail: ", result)
    }
}
