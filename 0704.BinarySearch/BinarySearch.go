func search(nums []int, target int) int {
    head, tail := -1, len(nums)
    var middle int
    for tail - head > 1 {
        middle = int((tail+head)/2)
        if nums[middle] == target {
            return middle
        } else if nums[middle] < target {
            head = middle
        } else {
            tail = middle
        }
    }
    return -1
}