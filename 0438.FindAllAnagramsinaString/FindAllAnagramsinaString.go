package main

func FindAnagrams(s string, p string) []int {
    result := []int{}
	tmp := make(map[byte]int)
    for _, val := range p {
        tmp[byte(val)] += 1
    }
    for i:=0; i <= len(s)-len(p); i++ {
		isSame := true 
		copyTmp := make(map[byte]int)
		for key, val := range tmp {
			copyTmp[key] = val
		}
		for j:=i; j < i+len(p); j++ {
			copyTmp[s[j]]--
			if copyTmp[s[j]] < 0 {
				isSame = false
				break
			}
		}
		if isSame { 
			result = append(result, i) 
		}
    }
    return result
}

func main() {

}
