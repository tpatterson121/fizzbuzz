package main

import (
	"fmt"
	"strconv"
)
	

func fizz() {
	// for loop from 1 to 20 
	for i := 1; i < 20; i++ {
		fmt.Printf("%v\n", i)
		// in the range of number, if remainder of modulo 3 equals 0, print fizz
		if i%3 == 0 {
			fmt.Printf("Fizz")
		}
		// in the range of numbers, if remainder equals 0, print buzz
		if i%5 == 0 {
			fmt.Printf("Buzz")
		}
		// spaces between each number in list
		fmt.Printf("\n")
	}
}

//refined format after i gave it attempts but overall, 
//the modulo and for loops made sense


func palin() {
	var a bool

	s := "racecar" 
	a, _ = strconv.ParseBool(s)

	fmt.Println(a)

		
	}

