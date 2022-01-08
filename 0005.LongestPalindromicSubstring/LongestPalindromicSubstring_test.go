package main

import (
	"testing"
)

func TestLongestPalindromicSubstringSence1(t *testing.T) {
    output := LongestPalindrome("babad")
    if output == "bab" || output == "aba" {
        t.Log("success")
    } else {
        t.Error("fail")
    }
}

func TestLongestPalindromicSubstringSence2(t *testing.T) {
    output := LongestPalindrome("cbbd")
    if output == "bb" {
        t.Log("success")
    } else {
        t.Error("fail")
    }
}

func TestLongestPalindromicSubstringSence3(t *testing.T) {
    output := LongestPalindrome("bb")
    if output == "bb" {
        t.Log("success")
    } else {
        t.Error("fail")
    }
}

func TestLongestPalindromicSubstringSence4(t *testing.T) {
    output := LongestPalindrome("ccc")
    if output == "ccc" {
        t.Log("success")
    } else {
        t.Error("fail")
    }
}

func TestLongestPalindromicSubstringSence5(t *testing.T) {
    output := LongestPalindrome("aaaa")
    if output == "aaaa" {
        t.Log("success")
    } else {
        t.Error("fail")
    }
}
