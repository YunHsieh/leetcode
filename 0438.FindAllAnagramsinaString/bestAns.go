package main

/*
s: ababababababab
p: aab
[aba] == aab
 [bab] == aab
  [aba] == aab
    ...

---
computing way:
head arr[head]--
tail arr[tail]++

start:
[1+1 1] == [2 1]
[1-1 1+1] == [2 1]
[1+1 1-1] == [2 1]

*/

func findAnagrams(s string, p string) []int {
    sArr, pArr := [26]int{}, [26]int{}
    ans := []int{}
    for _, v := range p {
        pArr[v - 'a']++
    }
    for i, v := range s {
        if i >= len(p) { 
            sArr[s[i - len(p)] - 'a']-- 
        }
        sArr[v - 'a']++
        if sArr == pArr { 
            ans = append(ans, i - len(p) + 1) 
        }
    }
    return ans
}