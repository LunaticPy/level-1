package t13

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

/*
4
0
1
2
3
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc000014138)
        /usr/lib/go/src/runtime/sema.go:56 +0x45
sync.(*WaitGroup).Wait(0xc000014130)
        /usr/lib/go/src/sync/waitgroup.go:130 +0x65
main.main()
        /home/lunatic/work/level1/13.go:17 +0xbb
exit status 2


так в функцию передается копия waitgroup, не связанная с остальными рутинами

*/
