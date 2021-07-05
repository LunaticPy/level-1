package t33

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	data := make(chan int)
	rez := make(chan int)
	end := make(chan bool)

	go func(data *chan int, rez *chan int, end *chan bool) {
		var buf int
		for {
			select {
			case buf = <-(*data):
				if buf%2 == 0 { // проверяем четность
					*rez <- buf
				}
			case <-*end:
				fmt.Println("end")
				return
			}
		}
	}(&data, &rez, &end)
	go func(data *chan int, rez *chan int, end *chan bool) {
		var buf int
		for {
			select {
			case buf = <-(*rez):
				fmt.Println(buf) // пишем в stdout

			case <-*end:
				fmt.Println("end")
				return
			}
		}
	}(&data, &rez, &end)

	for i := 0; i < 10; i++ {
		data <- rand.Int() // пишем в канал
	}

	time.Sleep(time.Second * 5)
	end <- true
	end <- true
	time.Sleep(time.Second)
}
