package main

import (
	"fmt"
	"os"
)

func main() {
	x := os.Args[1:]
	for _, y := range x {
		if y == "01" || y == "galaxy" || y == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}