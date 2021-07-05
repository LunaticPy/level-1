package t28

import "fmt"

func Adapter(a int64) int {
	return int(a)
}

func main() {
	a := int64(1)
	fmt.Println(Adapter(a) + int(2))
}
