package t21

import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	r  []int
	mt sync.RWMutex
}

func main() {
	data := Container{
		r: []int{1, 2, 3, 4, 5, 6, 7},
	} //делаем контейнер с мьютексом, который позвоялет одновременно читать из массива

	for i := 0; i < 5; i++ {
		go func(data *Container) {

			for i := 0; i < len(data.r); i++ {
				data.mt.RLock()
				fmt.Print(data.r[i])
				data.mt.RUnlock()
			}
			fmt.Println()

		}(&data)
	}
	time.Sleep(time.Second)
}
