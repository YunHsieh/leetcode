package main

/**
https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/discuss/812149/Thought-Process-How-Binary-Search-comes-to-mind
**/


func AnswerShipWithinDays(weights []int, D int) int {
    sum := 0
    for _,v :=range weights{
        sum += v
    }
    l, r := 1,sum
    for l < r{
        mid := l + (r-l)/2
        if valid(weights,D,mid){
            r = mid
        }else{
            l = mid+1
        }
    }
    return r
    
}

func valid(weights []int, D int, mid int) bool{
    tmpSum, count := 0, 0
    for i,v := range weights{
        if tmpSum + v <= mid{
            tmpSum += v
            continue
        }else{
            if v > mid {return false}
            tmpSum = v
            count ++
            if count == D && i != len(weights){
                return false
            }
        }
    }
    return true
}