/*
 If it is a multiple of 3 or 5 that will show FizzBuzz
 If it is a multiple of 5 that will show  Buzz
 If it is a multiple of 3 that will show  Fizz
 else print number
*/

package main

import (
	"fmt"
)

func fizzBuzz(n int32) {
    for i:=int32(1); n>=i; i++{
        if i % 3 == 0 && i % 5 == 0{
            fmt.Println("FizzBuzz")
        } else if i % 5 == 0{
            fmt.Println("Buzz")
        } else if i % 3 == 0{
            fmt.Println("Fizz")
        } else {
            fmt.Println(i)
        }
    }
}

func main(){
	fizzBuzz(15)
}
