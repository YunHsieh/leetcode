package main

import (
	"testing"
)

func TestValidParenthesesScence1(t *testing.T) {
    result := IsValid("([)]")

    if result == false {
        t.Log("> TestValidParenthesesScence1 success")
    } else {
        t.Error("fail")
    }
}

func TestValidParenthesesScence2(t *testing.T) {
    result := IsValid("()[]{}")

    if result == true {
        t.Log("> TestValidParenthesesScence2 success")
    } else {
        t.Error("fail")
    }
}

func TestValidParenthesesScence3(t *testing.T) {
    result := IsValid("[")

    if result == false {
        t.Log("> TestValidParenthesesScence3 success")
    } else {
        t.Error("fail")
    }
}

func TestValidParenthesesScence4(t *testing.T) {
    result := IsValid("]")

    if result == false {
        t.Log("> TestValidParenthesesScence4 success")
    } else {
        t.Error("fail")
    }
}
