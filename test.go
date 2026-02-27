package main

//program will not run/compile if any variables or libraries are unused
//go will infer variables but can be specified.
//print statements seem to auto append spaces between added varialbes. This doesn't occur when adding strings.
//go has no ternary if statements. JS ternary operator is '?'
import (
	"fmt"
	"time"
	//"math"
)

const chungus = "big chungus"

// types are declared AFTER a variable
func adder(x int, y int) int {
	return x + y
}

func main() {
	//fmt.Println("Hello, World!")
	//fmt.Println("go"+"lang", 7.2/3.1, 1+1, "Yep. That is", true, "and not", false)
	//var a = "This is the beginning"
	//var g int
	//g := 3
	//g = 4
	//goofy := "Potato Tomato"
	//fmt.Println(a, g, goofy)
	//i := 1
	fmt.Println("Big" + "Chungus")
	//basic for loop conditional
	//for i <= 3 {
	//	fmt.Println(i)
	//	i += 1
	//}
	//c-styled for-loop
	//for j := 0; j < 3; j++ {
	//	fmt.Println(j)
	//}
	//python style for loop
	//for j := range 3 {
	//	fmt.Println(j)
	//}
	//non-descript for loop (seems like a while true loop, no end conditional)
	//for {
	//	fmt.Println("looping!")
	//	break
	//}
	fmt.Println("No more looping.")
	//python style for loop with inner conditional
	for n := range 16 {
		if n%2 == 0 && n%3 == 0 {
			fmt.Println(n, "- This number is divisible by 2 and 3!")
		} else if n%3 == 0 {
			fmt.Println(n, "- this one is divisible by 3!")
		} else {
			fmt.Println(n)
		}
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend ~")
	default:
		fmt.Println("Weekday :[")
	}

	t := time.Now()
	fmt.Println("Current time", t)

	//switches without an expression can be used
	//as an alternate to if/else branches
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning Time.")
	default:
		fmt.Println("Evening time.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool(ean)")
		case int:
			fmt.Println("I am an integer.")
		default:
			fmt.Printf("I'm something else - %T\n", t)
		}
	}
	whatAmI(true)            //this will hit the bool case in our switch
	whatAmI(1)               //this will hit the int case in our switch
	whatAmI("Hello, World!") //%T and .(type) can be used to reference the type of a variable
	fmt.Println(adder(100, 200))
}
