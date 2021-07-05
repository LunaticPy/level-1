package t32

import (
	"fmt"
	"time"
)

func MySleep(slp int32) {

	dr := time.Duration(slp) * time.Second
	start := time.Now().Add(dr)
	for start.After(time.Now()) { // просто цикл проверяющий время
	}

	// или

	timer1 := time.NewTimer(dr)
	<-timer1.C // блокирует рутину до истечения времени

}

func main() {
	fmt.Println("Start")
	MySleep(2)
	fmt.Println("Done")
}
