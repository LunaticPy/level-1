package t22

import "fmt"

func main() {
	data := []int{12, 22, 333, 2, 21, 444, 1, 2, 5}
	QuickSort(data)

	fmt.Println(data)
}

func QuickSort(data []int) {
	core := len(data) - 1 // берем самый правый элемент, как опорный
	for i := 0; i < core; i++ {
		if data[i] >= data[core] {
			shift := 1
			if data[i] == data[core] {
				shift--
			}
			buf := data[i]
			sl := data[core+shift:]
			data = append(data[:i], data[i+1:core+shift]...)
			data = append(data, buf)
			data = append(data, sl...)
			core--
			i--
		}
	}

	if len(data) > 2 {

		QuickSort(data[:core])
		QuickSort(data[core:])
	}

}
