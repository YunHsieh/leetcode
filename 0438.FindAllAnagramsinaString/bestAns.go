package main

func findAnagrams(s string, p string) []int {
    sArr, pArr := [26]int{}, [26]int{}
    ans := []int{}
    for _, v := range p {
        pArr[v - 'a']++
    }
    for i, v := range s {
        if i >= len(p) { sArr[s[i - len(p)] - 'a']-- }
        sArr[v - 'a']++
        if sArr == pArr { ans = append(ans, i - len(p) + 1) }
    }
    return ans
}