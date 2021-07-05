package t4

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//инициализация
	var num int
	fmt.Println("Enter num of workers: ")
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //перехват сигналов
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic(err)
	}

	data := []chan int{} //создание списка каналов
	for i := 0; i < num; i++ {
		data = append(data, make(chan int, 1))
	}

	for i := 0; i < num; i++ {
		go func(data chan int, done chan bool, num int) {
			var out int

			for {
				select {
				case out = <-data: //ловля данных
					fmt.Printf("Worker %d print %d \n", num, out)

				case <-done: //ловля команды завершения
					fmt.Printf("Worker %d done\n", num)
					return
				case <-sigs: //ловля сигнала
					return
				}
			}
		}(data[i], done, i)
	}

	for i := 0; i < num*3; i++ {
		data[i%num] <- i // отправка данных

	}

	for i := 0; i <= num; i++ {
		done <- true

	}

}
