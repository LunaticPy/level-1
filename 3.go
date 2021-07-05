package t3

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) //делаем ограничение, чтобы горутины конкурировали за ресурсы на уровне преключения контекста планировщиком
	data := []int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup

	for _, d := range data {
		wg.Add(1)
		go func(data int, sum *int, wg *sync.WaitGroup) {
			*sum += data * data // можно сделать и с мьютексами, как в следующих задачах
			wg.Done()

		}(d, &sum, &wg)
	}
	wg.Wait() // ждем завершения всех рутин

	fmt.Println(sum)
}
