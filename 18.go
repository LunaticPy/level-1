package t18

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}

/*

[100 2 3 4 5]


v[0] = 100  обращение по ссылке на данные в копии слайса
v = append(v, b) модификация копии слайса

*/
