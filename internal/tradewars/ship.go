package main

import "fmt"

type ship struct {
	// inventoy (map[key]value)
	water_tank float64
	currency uint64
	fuel_tank float64
	top_speed float64
	current_speed float64
}

// func insert (inventory, item) {
// 	inventory.append(item)
// }

func main(){
	a_ship := ship{water_tank: 75.0,
				   currency: 0,
				   fuel_tank: 150.0,
				   top_speed: 225.0,
				   current_speed: 0}
	fmt.Println(a_ship)
}