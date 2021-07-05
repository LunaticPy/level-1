package t10

import (
	"fmt"
	"math"
)

func main() {

	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // температура
	points := []float64{-20, 10, 20}                                     // группы
	step := 10.0                                                         // окно группы
	rez := make(map[float64][]float64)                                   //результат

	for _, p := range points {
		buf := make([]float64, 0)
		for _, d := range data {
			if p >= 0 && p < d {
				if p-d >= -step {
					buf = append(buf, d)
				}
			} else if p < 0 && p > d {
				if math.Abs(d-p) <= step {
					buf = append(buf, d)
				}
			}
			rez[p] = buf
		}
	}

	fmt.Println(rez)

}
