package test

import "fmt"

type Foo interface {
	qux()
}

type bar struct {

}

type baz struct {

}

func (a bar) qux()  {
	fmt.Println("qqqqqq")
}

func (b baz) qux(){
	fmt.Println("zzzzzz")
}

func (a *bar) getA(){
	fmt.Println("sssss")
}

func qq()  {
	var f Foo
	f = bar{}
	f.qux()
	a := bar{}
	a.getA()
	f = baz{}
	f.qux()


}