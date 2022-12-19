package main

func ContainsNearbyDuplicate(nums []int, k int) bool {
    mapping := map[int]int{}
    for i, val := range nums {
        if _, ok := mapping[val]; ok && i - mapping[val] <= k {
            return true
        }
        mapping[val] = i
    }
    return false
}

func main() {

}
