package t27

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	buf := strings.Split(s, " ")

	for i, j := 0, len(buf)-1; i < j; {
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return strings.Join(buf, " ")
}

func main() {
	fmt.Println(Reverse("snow dog sun"))
}
