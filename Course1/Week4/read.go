package main

import (
	"fmt"
	"os"
)

func main() {

	type Name struct {
		fname string
		lname string
	}

	names := make([]Name, 0, 3)

	var filename string
	fmt.Printf("Name of file? ")
	fmt.Scan(&filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		barr := make([]byte, 43) // 43 for windows line endings, 42 for unix
		nb, _ := f.Read(barr)
		// fmt.Print(barr)
		if nb < 43 {
			break
		}
		newfname := string(barr[0:20])
		// fmt.Print(newfname)
		newlname := string(barr[21:41])
		newname := Name{fname: newfname, lname: newlname}
		names = append(names, newname)
	}
	for _, v := range names {
		fmt.Printf("First name: %s, Last name: %s\n", v.fname, v.lname)
	}
	f.Close()
}
