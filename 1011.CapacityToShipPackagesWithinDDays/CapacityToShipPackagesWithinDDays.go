package main

import (
    . "fmt"
) 

func AbsoluteValue(num int) int {
    if num < 0 {
        return -1*num
    }
    return num
}

func ShipWithinDays(weights []int, days int) int {
    packages := make([][]int, days)
    packages[0] = weights
    summary := make([]int, days)

    for _, val := range weights {
        summary[0] += val
    }

    isBalance := true

    for isBalance {
        switch_count := 0
        for i:=1; len(packages) > i; i++ {
            last_element := packages[i-1][len(packages[i-1])-1]
            switch_spacing := AbsoluteValue((summary[i-1] - last_element) - (summary[i] + last_element))
            original_spacing := AbsoluteValue(summary[i-1] - summary[i])
            if switch_spacing <= original_spacing || len(packages[i]) == 0 {
                packages[i-1] = packages[i-1][:len(packages[i-1])-1]
                packages[i] = append([]int{last_element}, packages[i]... )
                summary[i-1] -= last_element
                summary[i] += last_element
                switch_count += 1
            }
        }
        if switch_count == 0 {
            isBalance = false
        }
    }

    max := 0
    for _, val := range summary {
        if val > max {
            max = val
        }
    }
    return max
}

func main () {
    max := ShipWithinDays([]int{10,5,2,4,8,9,1,8,1,2}, 4)
    Println(max)
}
