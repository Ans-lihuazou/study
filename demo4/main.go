package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3, 3)
	fmt.Println(slice)

	mp := make(map[string]string, 3)
	mp["SH"] = "ShangHai"
	mp["BJ"] = "BeiJing"
	fmt.Println(mp)
}
