/*
 If it is a multiple of 3 or 5 that will show FizzBuzz
 If it is a multiple of 5 that will show  Buzz
 If it is a multiple of 3 that will show  Fizz
 else print number
*/

package main

import (
    "strconv"
)

func fizzBuzz(n int64) string {
    result := ""
    for i:=int64(1); n>=i; i++{
        if i % 3 == 0 && i % 5 == 0{
            result += "FizzBuzz"
        } else if i % 5 == 0{
            result += "Buzz"
        } else if i % 3 == 0{
            result += "Fizz"
        } else {
            result += strconv.FormatInt(i, 10)
        }
    }
    return result
}
