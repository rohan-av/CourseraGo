package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f string
	fmt.Scan(&f)
	ff, _ := strconv.ParseFloat(f, 32)
	fmt.Printf(strconv.Itoa(int(ff)))
}
