package main

func isPossible(nums []int) bool {
    endsCount := map[int]int{}
    count := map[int]int{}
    
    for _, element := range nums {
        count[element]++
    }
    for _, element := range nums {
        if count[element] == 0 {
            continue
        }
        count[element]--
        if endsCount[element-1] > 0 {  
            endsCount[element-1]--
            endsCount[element]++
        } else if count[element+1] > 0 && count[element+2] > 0 {
            endsCount[element+2]++
            count[element+1]--
            count[element+2]--
        } else {
            return false
        }
    }
    return true
}

func main() {

}
