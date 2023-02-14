package main

import (
	"fmt"
	"main/telegraf"
)

func main() {
	Method1()
	telegraf.MethodTelegraf()
}
func Method1() {
	fmt.Println("This is the method from the main file")
}
