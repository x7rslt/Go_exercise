package method

import ("fmt";"testing")
//方法：Go支持在结构类型上定义的方法

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func TestMethod(t *testing.T) {
	var tangle rect = rect{
		width:  12,
		height: 15,
	}
	fmt.Println(tangle.area())
	fmt.Println(tangle.perim())

	rtangle := &tangle

	fmt.Println(rtangle.area())
	fmt.Println(rtangle.perim())

}
