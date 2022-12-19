package main

import (
    "strings"
)

func minRemoveToMakeValid(s string) string {
    results := strings.Split(s, "")
    stack := []byte{}
    for i:=0; i < len(results); i++ {
        if s[i] == '(' {
            stack = append(stack, s[i])
        } else if s[i] == ')' {
            if len(stack) != 0 {
                stack = stack[:len(stack)-1]
            } else {
                results[i] = ""
            }
        }
    }
    if len(stack) != 0{
        for i:=len(results)-1; i >= 0; i-- {
            if len(stack) !=0 {
                if results[i] == "(" {
                    results[i] = ""
                    stack = stack[:len(stack)-1]
                }
            } else {
                break
            }
        }
    }
    return strings.Join(results, "")
}

func main() {

}
