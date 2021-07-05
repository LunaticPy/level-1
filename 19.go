package main //t19

import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {

		v += "A"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	l := 100

	justString = v[:l]
	fmt.Println(len([]byte(justString)))
	fmt.Println(cap([]byte(justString)))

	b := []rune(v)[0:l:l]
	fmt.Println(len((b)))
	fmt.Println(cap((b)))

}

func main() {
	someFunc()
}

/*

при применении v[:100] мы будем таскать все 2^10 ячеек а не 100, для проверки можно взять sizeof или cap для slice
для исправления можно использовать связку make([]type, 100), copy(b, a), тогда мы будем брать только то, что нужно
у string поведение особое, оно будет тянуть запас по cap'у однако достаточно малое, но для надежности можно использовать этот же подход
так же можно использовать 3х индексный срез[0:l:l]


*/
