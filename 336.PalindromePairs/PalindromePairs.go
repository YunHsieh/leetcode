// It's not worked
package main

func isPalindromePairs(word string) bool{
    n := len(word)
    mid := int(n/2)
    h, t := mid-1, mid
    if n % 2 != 0 {
        h, t = mid-1, mid+1
    }
    for ; h>=0; {
        if word[h] != word[t] {
            return false
        }
        h--
        t++
    }
    return true
}

func palindromePairs(words []string) [][]int {
    results := make([][]int, 0)
    for i:=0; i<len(words); i++ {
        for j:=0; j<len(words); j++ {
            if i!=j && isPalindromePairs(words[i]+words[j]){
                results = append(results, []int{i, j})
            }
        }
    }
    return results
}

func main() {

}
