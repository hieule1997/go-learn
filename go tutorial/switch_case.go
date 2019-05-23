package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	// today := time.Now().Weekday()
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().Clock())
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().Nanosecond())
	// fmt.Println(time.Now().UnmarshalText())
	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("Today.")
	// case today + 1:
	// 	fmt.Println("Tomorrow.")
	// case today + 2:
	// 	fmt.Println("In two days.")
	// default:
	// 	fmt.Println("Too far away.", today)
	// }
}
