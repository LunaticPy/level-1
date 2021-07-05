package t2

import (
	"fmt"
	"runtime"
)

func main() {
	data := []float32{2, 4, 6, 8, 10}
	runtime.GOMAXPROCS(1) //делаем ограничение, чтобы горутины конкурировали за ресурсы на уровне преключения контекста планировщиком

	for _, d := range data {
		go func(data float32) {
			fmt.Println(data * data)

		}(d)
	}
}
