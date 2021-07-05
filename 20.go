package t20

import (
	"fmt"
	"time"
)

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice)
	}(slice)
	time.Sleep(time.Second)
	fmt.Println(slice)
}

/*

[b b a]

[a a]

делаем копию slice'а и забываем про нее после завершения функции


*/
