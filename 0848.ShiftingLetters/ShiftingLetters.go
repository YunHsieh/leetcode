/*
note: To be careful to overflow by the integer
*/

package main

func ShiftingLetters(s string, shifts []int) string {
	res := make([]rune, len(shifts))
	total := 0
    for i:=len(s)-1; i>=0; i-- {
		total += shifts[i] % 26
		res[i] = rune((int(s[i]) - 'a' + total)) % 26 + 'a'
    }
    return string(res)
}

func main() {

}
