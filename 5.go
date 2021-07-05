package t5

import (
	"fmt"
	"time"
)

func main() {
	N := 5
	data := make(chan int, 1)
	signal := make(chan bool, 1)
	//рутина отправки
	go func(data chan int, signal chan bool) {
		for {
			select {
			case <-signal: // ловля сигнала
				fmt.Printf("Writer done\n")
				return
			default: // отправка данных
				time.Sleep(1 * time.Second)
				data <- time.Now().Second()
			}
		}
	}(data, signal)
	//рутина получения
	go func(data chan int, signal chan bool) {
		var out int
		for {
			select {
			case <-signal: // ловля сигнала
				fmt.Printf("Reader done\n")
				return
			case out = <-data: // ловля данных
				fmt.Printf("Writer %d print\n", out)
			}
		}
	}(data, signal)

	time.Sleep(time.Duration(N) * time.Second) //таймер
	signal <- false
	signal <- false
	time.Sleep(time.Second)
	fmt.Printf("Done\n")
}
