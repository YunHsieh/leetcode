package main

import (
	. "fmt"
)

func comparingUntilNotEqual(s string, l int, r int) ([]int) {
    moveIndex := 0
    for ;(l-moveIndex) > 0 && (r+moveIndex) < len(s); {
        Println(l-moveIndex, r+moveIndex)
        if s[l-moveIndex] != s[r+moveIndex] {
            moveIndex--
            return []int{l-moveIndex, r+moveIndex}
        }
        moveIndex++
    }
    
    if (l-moveIndex) < 0 || (r+moveIndex) >= len(s) || s[l-moveIndex] != s[r+moveIndex] {
        moveIndex--
    }
    return []int{l-moveIndex, r+moveIndex} 
}

func LongestPalindrome(s string) string {
    result := []int{0,0}
    compared := result
    if len(s) <= 1 {
        return s
    }
    
    for i := 1; i < len(s); i++ {
        if s[i-1] == s[i] {
            compared = comparingUntilNotEqual(s, i-1, i)
        }
        if (result[1] - result[0]) < (compared[1] - compared[0]) {
            result = compared
        }
        if i < len(s)-1 && s[i-1] == s[i+1] {
            compared = comparingUntilNotEqual(s, i, i)
        }
        if (result[1] - result[0]) < (compared[1] - compared[0]) {
            result = compared
        }
    }
    return s[result[0]:result[1]+1]
}
