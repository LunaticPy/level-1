package t30

func main() {
	num := 2
	data := []int{1, 2, 2, 5, 12, 21, 22, 333, 444}
	// ======
	data = append(data[:num], data[num+1:]...)

	// или

	copy(data[num:], data[mun+1:])
	// ======

	data = data[:len(data)-1]
}
