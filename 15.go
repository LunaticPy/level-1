package t15

import "fmt"

func main() {
	a, b := 1, 2

	fmt.Println(a, b)

	a, b = b, a // не понятно, что еще тут требовалось
	fmt.Println(a, b)
}
