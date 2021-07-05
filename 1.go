package t1

type Human struct {
	Age  int
	Name string
}

func (h Human) do() {
	a := 1 + 1
}

type Action struct {
	Human Human
	Name  string
}

func (a Action) do() {
	a.Human.do()

}
