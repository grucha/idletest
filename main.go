package main

import "fmt"
import "./os_specific"

func main() {
	fmt.Println("Hello")
	os_specific.ShowInactiveTime()
}
