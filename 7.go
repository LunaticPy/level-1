package t7

// https://habr.com/ru/post/338718/
import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	data := []float32{2, 4, 6, 8, 10}
	rez := make(map[float32]float32)
	runtime.GOMAXPROCS(1) //делаем ограничение, чтобы горутины конкурировали за ресурсы на уровне преключения контекста планировщиком
	var wg sync.WaitGroup

	for _, d := range data {
		wg.Add(1)
		go func(data float32, wg *sync.WaitGroup) { //альтернатива - mutex
			rez[data] = data
			fmt.Println(rez)
			wg.Done()

		}(d, &wg)
	}
	wg.Wait()

}
