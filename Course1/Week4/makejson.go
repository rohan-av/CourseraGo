package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var m map[string]string
	m = make(map[string]string)
	var name string
	var addr string
	fmt.Printf("Name? ")
	fmt.Scan(&name)
	fmt.Printf("Address? ")
	fmt.Scan(&addr)
	m["name"] = name
	m["address"] = addr
	barr, _ := json.Marshal(m)
	fmt.Printf(string(barr))
}
