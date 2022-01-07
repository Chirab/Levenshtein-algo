package main

import (
	"fmt"
	"sort"
)


func  minVal(a, b, c int) int {
	test := []int{a, b, c}
	sort.Ints(test)
	return test[0]
}

func leven(s, d string) int{
	sl := len(s)
	dl := len(d)
	sarr := make([][]int, sl + 1)

	//initialize 2d slice
	for i := range sarr {
		sarr[i] = make([]int,  dl + 1)
	}

	//fil the first column and row
	for i := range sarr[0] {
		sarr[0][i] = i
	}
	for j := range sarr {
		sarr[j][0] = j
	}

	for i:= 1; i < len(s) + 1; i++ {
		for j:= 1; j < len(d) + 1; j++ {
			if d[j - 1] == s[i - 1] {
				sarr[i][j] = sarr[i - 1][j - 1]
			} else {
				min := minVal(sarr[i - 1][j - 1], sarr[i - 1][j], sarr[i][j - 1])
				sarr[i][j] = min + 1
			}
		}
	}
	return sarr[sl][dl]
}

func main() {
	pdt := leven("20 Rue Léescueyer", "20 Rue Lécuyer")
	fmt.Println(pdt)
}
