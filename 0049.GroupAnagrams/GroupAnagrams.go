package main

func groupAnagrams(strs []string) [][]string {
    words := make(map[string][]string)
    results := [][]string{}
    for _, val := range strs {
        sList := strings.Split(val, "")
        sort.Strings(sList)
        strSorted := strings.Join(sList, "")
        words[strSorted] = append(words[strSorted], val)
    }
    for _, val := range words {
        results = append(results, val)
    }
    return results
}
