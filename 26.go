package t26

import "fmt"

func Reverse(s string) string {
	buf := []rune(s)

	for i, j := 0, len(buf)-1; i < j; {
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func main() {
	fmt.Println(Reverse("pppppppplllllllll"))
}
