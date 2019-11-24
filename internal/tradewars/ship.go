package main

import "fmt"

type ship struct {
	fuel_tank float64
	top_speed float64
}

// func insert (inventory, item) {
// 	inventory.append(item)
// }

func main(){
	a_ship := ship{fuel_tank: 150.0,
				   top_speed: 225.0}
	fmt.Println(a_ship.fuel_tank)
}