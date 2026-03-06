package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i , " as ")
	switch i {
		case 1 : 
			fmt.Println("one")
		case 2 :
			fmt.Println("two")
		case 3 : 
			fmt.Println("three")
	}

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday: 
			fmt.Println("Its the weekend")
		default:
			fmt.Println("Its a weekday")
	}

	t := time.Now()
	switch {
		case t.Hour() < 12 :
			fmt.Println("It's before noon")
		default : 
			fmt.Println("It's after noon")
	}

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
			case bool :
				fmt.Println("I'm a bool")
			case int :
				fmt.Println("I'm an int")
			default:
				fmt.Printf("Don't know the type  %T\n", t)
		}
	}

	whoAmI(true)
	whoAmI(1)
	whoAmI("hey")
}
