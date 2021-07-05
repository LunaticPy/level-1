package t23

import "fmt"

func main() {
	data := []int{1, 2, 2, 5, 12, 21, 22, 333, 444}
	rez := new(int)
	BinSearch(data, 5, rez)
	fmt.Println(*rez)
}

func BinSearch(data []int, val int, rez *int) bool {
	/*
		data - слайс в котором идет поиск
		val - значение для поиска
		rez - финальный индекс*/

	if len(data) > 2 {
		if BinSearch(data[:2], val, rez) {
			return true
		}
		if BinSearch(data[2:], val, rez) {
			*rez += 2
			return true
		}

	} else {
		if data[0] == val {
			*rez = 0
			return true
		} else if data[len(data)-1] == val {
			*rez = len(data) - 1
			return true
		}
	}
	return false
}
