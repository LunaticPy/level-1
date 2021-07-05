package t17

import "fmt"

func main() {

	var i interface{}

	i = int32(25)

	switch i.(type) { // проверка типов, можно еще много проверок сделать
	case int:
		fmt.Println("int", i.(int))
	case float32:
		fmt.Println("float32", i.(float32))
	case int64:
		fmt.Println("int64", i.(int64))
	case int32:
		fmt.Println("int32", i.(int32))
	case float64:
		fmt.Println("float64", i.(float64))
	default:
		fmt.Println("Unknoun")
		fmt.Println("И так далеее")

	}

}
