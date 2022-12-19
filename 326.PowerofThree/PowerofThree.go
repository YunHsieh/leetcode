package main

func isPowerOfThree(n int) bool {
    if n == 0 {
        return false
    } else if n == 1{
        return true
    }
    
    if n := float64(n) / 3; n == float64(int(n)) {
        return isPowerOfThree(int(n))
    } else {
        return false
    }
}
