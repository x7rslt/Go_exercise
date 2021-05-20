package test

import (
	"fmt"
	"math"
	"testing"
)

//变量的地址叫做指针,用&表示指针、*表示值
func TestPointer(t *testing.T) {
	var x int
	x = 7

	p := &x
	value := *p

	fmt.Println(x, p, *p, value)
}

type x int

func (i *x) Set() {
	*i = 10
}

func TestPointerFunc(t *testing.T) {
	var j x
	j.Set()
	fmt.Println(j)
}

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
func TestNew(t *testing.T) {
	//var p *Point   //panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	//修改：
	p := new(Point)
	fmt.Println(p.Abs())
	p2 := Point{10, 10}
	fmt.Println(p2.Abs())
}
