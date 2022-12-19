package main

func IsValid(s string) bool {
    symbols := map[byte]byte{
        '(': ')',
        '[': ']',
        '{': '}',
    }
    stack := []byte{}
    for i:=0; len(s)>i; i++ {
        if val, ok := symbols[s[i]]; ok {
            stack = append(stack, val)
        } else {
            if len(stack) == 0 || stack[len(stack)-1] != s[i] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    if len(stack) != 0 {
        return false
    }
    return true
}

func main() {

}
