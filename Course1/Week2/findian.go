package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter string here: ")
	s, _ := reader.ReadString('\n')
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "i") && strings.Contains(s, "a") && strings.HasSuffix(s, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
