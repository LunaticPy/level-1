package t12

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}

/*
выдаст 1 1

тк значене b локально то поинтер при таком присвоении, примет локальное значение, которое потом умрет,
чтобы p изменилось надо написать *p = ....

*/
