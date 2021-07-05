package t11

import (
	"fmt"
	"sync"
	"time"
)

type Counters struct {
	mx sync.Mutex
	m  map[int]bool
}

func main() {
	data := []int{19, 15, 24, -21, 32, -25, -27, 13}
	data2 := []int{2, 19, 4, 6, 8, -27}
	rez := &Counters{m: make(map[int]bool)} // конкурентная запись в map с помощью Mutex

	for _, d := range data {
		go func(d int, rez *Counters) {
			for _, d2 := range data2 {
				if d2 == d {
					rez.mx.Lock()
					rez.m[d2] = true
					rez.mx.Unlock()
				}
			}
		}(d, rez)
	}
	time.Sleep(time.Microsecond * 5)
	fmt.Println(rez.m)

}
