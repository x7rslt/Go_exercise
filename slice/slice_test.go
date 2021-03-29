package slice_test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1)
	a[2] = 42
	fmt.Println(a, b)

	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99
	fmt.Println(foo, bar)
}

func TestSliceCap(t *testing.T) {
	path := []byte("AAAA/BBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex:]

	fmt.Println("dir1:", string(dir1))
	fmt.Println("dir2:", string(dir2))

	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1:", string(dir1))
	fmt.Println("dir2:", string(dir2))

	//修改内存覆盖,使用Full Slice Expression
	path2 := []byte("AAAA/BBBBBBBB")
	sepIndex2 := bytes.IndexByte(path2, '/')
	Newdir1 := path2[:sepIndex2:sepIndex2]
	Newdir2 := path2[sepIndex2:]

	fmt.Println("Newdir1:", string(Newdir1))
	fmt.Println("Newdir2:", string(Newdir2))

	Newdir1 = append(Newdir1, "suffix"...)

	fmt.Println("Newdir1:", string(Newdir1))
	fmt.Println("Newdir2:", string(Newdir2))
}

func TestReflect(t *testing.T) {
	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	//prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true
}

type Person struct {
	Name   string
	Sexual string
	Age    int
}

func PrintPerson(p *Person) {
	fmt.Println(p.Name, p.Sexual, p.Age)
}

func (p *Person) Print() {
	fmt.Println(p.Name, p.Age, p.Sexual)
}
//在 Go 语言中，使用“成员函数”的方式叫“Receiver”，这种方式是一种封装，因为 PrintPerson()本来就是和 Person强耦合的，所以，理应放在一起。
func TestStruct(t *testing.T) {
	var p = Person{
		Name:   "Xiao",
		Sexual: "Male",
		Age:    12,
	}

	PrintPerson(&p)
	p.Print()
}
