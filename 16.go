package t16

import "fmt"

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}

/*

0

в локальном поле создается новая переменная, которая уничтожается после завершения if

*/
