package main

import (
	"fmt"

	"github.com/mrcoet/quanbit_golang/work"
)

func main() {
	// fmt.Println(work.DropRock(4, "0"))
	// ranHex := work.DropRock(4, "0")
	// ranHex := "0000000000000000000000000000000000000000000000000000000000000001"
	// work.CheckHex(ranHex)
	i := 1
	for i <= 10000 {
		fmt.Println(i)
		work.CheckHex(work.DropRock(4, "0"))
		i++
	}
}
