package t9

import (
	"fmt"
)

func main() {
	raw := []int{2, 4, 6, 8, 10}
	data := make(chan int)
	rez := make(chan int)
	sig := make(chan bool)

	go func(data chan int, rez chan int) {
		var out int
		for {
			select {
			case out = <-data: // обработка данных
				rez <- out * 2
			case <-sig:
				return
			}

		}
	}(data, rez)

	for _, r := range raw {
		data <- r                   // запись данных
		fmt.Println("Rez: ", <-rez) // чтение результата
	}
	sig <- true // завершение работы

	fmt.Printf("Done\n")
}
