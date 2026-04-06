package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Print("Enter you'r name : ")
	fmt.Scanln(&name)
	fmt.Print("Enter you'r age : ")

	_,err := fmt.Scanln(&age)

	if err != nil {
		fmt.Println("error",err)
		return
	}
	
	fmt.Printf("Hey %s! kudos to turning %d in hits year",name,age+1)
}