package recursion

import("fmt";"testing")

//递归Recursion
func Sum(i int) int {
	//fmt.Println(i)
	if i < 1000000 {
		i = i * 10
		fmt.Println(i)
		return Sum(i)
	}
	return 999999
}

func TestRecursion(t *testing.T) {
	fmt.Println(Sum(1))
}
