package main

import ("fmt"
		"math/rand")
func foo()  {
	fmt.Println("I am for beginner",rand.Intn(100))
}
func add(x,y float64) float64{
	return x+y
}
type car struct{
	gas_pedal uint16
	brake_pedal uint16
	top_speed_km float64
}

func main(){
	my_car := car{12,4,5.4}
	fmt.Println(my_car.gas_pedal)
}