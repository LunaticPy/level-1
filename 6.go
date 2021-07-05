package t6

import (
	"fmt"
	"os"
	"time"
)

/*Прервать горутину снаружи нельзя.

Горутина сама контролирует свое выполнение
Для планового завершения требуется регулярная проверка флагов посредством канала для общения извне/плановое завершение.


Исключение - os.Exit(...), но тогда останавливается вся программа сразу, никакие defer, финализаторы и т.п. не вызываются.*/
func main() {
	sig := make(chan bool)
	go SelfDone()
	time.Sleep(time.Second)
	go SignalDone(sig)
	sig <- false
	time.Sleep(time.Second)
	go EmergencyDone(sig)
	time.Sleep(time.Second)
	os.Exit(2)
}

func SelfDone() {
	fmt.Println("Done1")

}
func SignalDone(s chan bool) {
	select {
	case <-s:
		fmt.Println("Done2")
		return
	}

}
func EmergencyDone(s chan bool) {
	select {
	case <-s:
		fmt.Println("Done3")
		return
	}
}
