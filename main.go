package main

import "fmt"

// MyType blablabla
type MyType struct{ i int }

// Get1 blablabla
func (p *MyType) Get1() int {
	p.i = p.i + 1
	return p.i
}

// Get2 blablabla
func (p MyType) Get2() int {
	p.i = p.i + 1
	return p.i
}

func main() {
	var pm1 = new(MyType)
	var n1 = pm1.Get1()
	fmt.Println(n1)

	n1 = pm1.Get1()
	fmt.Println(n1)

	n1 = pm1.Get1()
	fmt.Println(n1)

	var pm2 = &MyType{0}
	var n2 = pm2.Get2()
	fmt.Println(n2)

	n2 = pm2.Get2()
	fmt.Println(n2)
}
