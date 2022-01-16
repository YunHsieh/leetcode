package main

import (
	"testing"
)

func TestShiftingLettersScence1(t *testing.T) {
    result := ShiftingLetters("abc", []int{3,5,9})
	ans := "rpl"
    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}


func TestShiftingLettersScence2(t *testing.T) {
    result := ShiftingLetters("aaa", []int{1,2,3})
	ans := "gfd"
    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}

func TestShiftingLettersScence3(t *testing.T) {
    result := ShiftingLetters("mkgfzkkuxownxvfvxasy", []int{505870226,437526072,266740649,224336793,532917782,311122363,567754492,595798950,81520022,684110326,137742843,275267355,856903962,148291585,919054234,467541837,622939912,116899933,983296461,536563513	})
	ans := "wqqwlcjnkphhsyvrkdod"
    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}
