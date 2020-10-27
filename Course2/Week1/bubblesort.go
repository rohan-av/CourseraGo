package main

import "fmt"

func Swap(sli []int, index int) {
	if index+2 <= len(sli) {
		tmp := sli[index]
		sli[index] = sli[index+1]
		sli[index+1] = tmp
	}
}

func BubbleSort(sli []int) {
	for i := range sli {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func PrintSlice(sli []int) {
	for i, v := range sli {
		fmt.Printf("%d", v)
		if i != len(sli)-1 {
			fmt.Printf(" ")
		} else {
			fmt.Printf("\n")
		}
	}
}

func main() {
	var n int
	fmt.Printf("Number of items? ")
	fmt.Scan(&n)
	if n > 10 {
		fmt.Printf("Number of items must be less than 10!")
	} else {
		sli := make([]int, 0, 10)
		for i := 0; i < n; i++ {
			var x int
			fmt.Scan(&x)
			sli = append(sli, x)
		}
		BubbleSort(sli)
		PrintSlice(sli)
	}
}
