package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    n, result := len(nums1), 0
	hashMap := map[int]int{}

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			hashMap[nums1[i] + nums2[j]] += 1
		}
	}

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			result += hashMap[0 - nums3[i] - nums4[j]]
		}
	}
	return result
}
