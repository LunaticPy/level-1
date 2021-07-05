package main //t8

import (
	"fmt"
)

func main() {
	pos, val, number := 6, 0, int64(64)

	if val == 1 {
		fmt.Println(number | 1<<pos) // логическое или - делаем pos 1
	} else {
		fmt.Println(number &^ (1 << pos)) // XOR без "не", то что спарава 1 будет обнулено, то что 0 скопируется из number
	}

}
