package t29

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b, z big.Float
	a.SetInt64(3222e10)
	b.SetInt64(3222e14)
	z.Add(&a, &b) // сложение
	fmt.Printf("z = %.10g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())
	z.Sub(&a, &b) // вычитание
	fmt.Printf("z = %.10g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())
	z.Mul(&a, &b) // умножение
	fmt.Printf("z = %.10g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())
	z.Quo(&a, &b) // деление
	fmt.Printf("z = %.10g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())
}
