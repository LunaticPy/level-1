package main //t25

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Container struct {
	counter int
	mx      sync.Mutex
}

func main() {
	num := 10
	data := Container{}
	end := make(chan bool)

	for i := 0; i < num; i++ {
		go func(data *Container, end *chan bool) {
			for i := 0; i < rand.Int(); i++ { //может быть 1 значение или несколько запустите несколько раз
				select {
				case <-*end:
					return
				default:
					data.mx.Lock()
					data.counter++
					fmt.Print(data.counter, " ")
					data.mx.Unlock()
				}
			}

		}(&data, &end)
	}
	time.Sleep(time.Microsecond * 15)
	end <- true
}
