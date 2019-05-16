package main

import "fmt"


func main()
{
	grades := make(map[string ]float32)
	grades["daica"] = 2
	grades["laodai"] = 20
	grades["daigia"] = 100

	fmt.Println(grades)
	delete(grades,"daica")
	fmt.Println(grades)
	for k , v := range grades{
		fmt.Println(k, " : ",v)
	} 

}