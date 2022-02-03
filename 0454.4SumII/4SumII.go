package main

func sumNum(i, j, k, l int) int {
    if i+j+k+l == 0 {
        return 1
    }
    return 0
}

func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    count := 0
    var arrCount [4]int
    for {
        count += sumNum(nums1[arrCount[3]], nums2[arrCount[2]], nums3[arrCount[1]], nums4[arrCount[0]])
        for i:=0; i<4; i++{
            if arrCount[i] >= len(nums1)-1 && i!=3 {
                arrCount[i] = 0
            } else {
                arrCount[i]++
                break
            }
        }
        if arrCount[3] >= len(nums1) {
            break
        } 
    }
    return count
}
