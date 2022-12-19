package main

func ValidateStackSequences(pushed []int, popped []int) bool {
	results := []int{}
	pushedN := len(pushed)
	pushedIndex := 0
	poppedIndex := 0
	n := len(results)
    for pushedN > pushedIndex || len(popped) > poppedIndex {
    	n = len(results)
    	if n > 0 && results[n-1] == popped[poppedIndex] {
    		results = results[:n-1]
    		poppedIndex++
    	} else if pushedN > pushedIndex {
    		results = append(results, pushed[pushedIndex])
    		pushedIndex++
    	} else {
    		break
    	}
    }
    if len(results) == 0 {
    	return true
    }
    return false
}

func main() {

}
