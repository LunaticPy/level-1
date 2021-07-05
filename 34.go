package t34

import (
	"fmt"
)

func main() {
	data := "asdfghjkl"
	rez := make(map[rune]int) // множество с подсчетом встреченных элементов
	flag := true

	for _, d := range data {
		if _, ok := rez[d]; !ok {
			rez[d] = 1
		} else {

			rez[d] += 1
			flag = false
		}

	}

	if flag {
		fmt.Println("Uniqe")
	} else {
		fmt.Println("Not Uniqe")
	}

}
