/*
 If it is a multiple of 3 or 5 that will show FizzBuzz
 If it is a multiple of 5 that will show  Buzz
 If it is a multiple of 3 that will show  Fizz
 else print number
*/

package main

import "testing"

func TestKthLargestElementinanArrayScence1(t *testing.T) {
    result := FindKthLargest([]int{3,2,1,5,6,4}, 2)
	ans := 5
    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}

func TestKthLargestElementinanArrayScence2(t *testing.T) {
    result := FindKthLargest([]int{1}, 1)
	ans := 1
    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}

func TestKthLargestElementinanArrayScence3(t *testing.T) {
    result := FindKthLargest([]int{2,1}, 2)
	ans := 1
    if result == ans {
        t.Log("success")
    } else {
        t.Error("fail", result)
    }
}