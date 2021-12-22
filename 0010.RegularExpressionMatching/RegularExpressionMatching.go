package main

import (
    . "fmt"
)

func isMatch(s string, p string) bool {
    s_index := 0
    for i := 0; i < len(p); i++ {
        if s[s_index] == '*' {
            continue
        }
        if s[s_index] == p[i] || p[i] == '.' {
            if len(p)-1 == i {
                return false
            }
            if p[i+1] == '*' {
                for ;; {
                    Println(s_index)
                    if len(s)-1 == s_index {
                        return true
                    }
                    if s[s_index] == p[i] || p[i] == '.' {
                        s_index += 1
                    } else {
                        i+=2
                        break
                    }
                }
            }
        }
        s_index += 1
        if len(s)-1 == s_index {
            return true
        }
        
    }
    return false
}

func main () {
    a := isMatch("ab", ".*c")
    Println(a)

}