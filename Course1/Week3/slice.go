package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var sli []int
	sli = make([]int, 0, 3)
	for {
		var s string
		fmt.Scan(&s)
		if s == "X" {
			break
		}
		d, _ := strconv.Atoi(s)
		sli = append(sli, d)
		sort.Ints(sli)
		for _, v := range sli {
			fmt.Printf("%d ", v)
		}
		fmt.Printf("\n")
	}
}
