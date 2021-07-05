package t14

import (
	"fmt"
)

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	rez := make(map[string]struct{}) // используем ключи, как множество

	for _, d := range data {
		rez[d] = struct{}{}
	}
	fmt.Println(rez)
}
