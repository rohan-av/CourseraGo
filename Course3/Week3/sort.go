package main

import (
	"fmt"
	"sync"
)

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

func SortSection(wg *sync.WaitGroup, sli []int, startIndex int, endIndex int) {
	subsli := make([]int, 0, endIndex-startIndex+1)
	for i := startIndex; i <= endIndex; i++ {
		subsli = append(subsli, sli[i])
	}
	BubbleSort(subsli)
	for i := startIndex; i <= endIndex; i++ {
		sli[i] = subsli[i-startIndex]
	}
	PrintSlice(sli)
	// PrintSlice(sli[startIndex:endIndex])
	wg.Done()
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
	var wg sync.WaitGroup
	fmt.Printf("How many elements? ")
	fmt.Scan(&n)
	subSize := n / 4
	sli := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		sli = append(sli, x)
	}
	wg.Add(4)
	go SortSection(&wg, sli, 0, subSize-1)
	go SortSection(&wg, sli, subSize, 2*subSize-1)
	go SortSection(&wg, sli, 2*subSize, 3*subSize-1)
	go SortSection(&wg, sli, 3*subSize, n-1)
	wg.Wait()
	BubbleSort(sli)
	PrintSlice(sli)
}
