package main

import (
	"fmt"
)

func main() {
    greet()
}


func reverse() {
	var slice = make([]int,6,10)
    fmt.Print(slice)
	fmt.Println()
	for i := 0; i < len(slice)/2; i++ {
        temp := slice[i]
		slice[i] = slice[len(slice)-1-i]
		slice[len(slice)-1-i] = temp
	}
	fmt.Print(slice, len(slice), cap(slice))
}

func even_odd() {
	var num int
    fmt.Print("Enter the number : ")
	fmt.Scanln(&num)
	if num % 2 == 0 {
		fmt.Printf("%d is even",num)
	}else {
		fmt.Printf("%d is odd",num)
	}
}

func greet() {
	var name string
    fmt.Print("Enter your name : ")
	fmt.Scanln(&name)
	var age int
    fmt.Print("Enter your age : ")
	fmt.Scanln(&age)
	
	fmt.Printf("Hey %s! kudos to turning %d in hits year",name,age+1)
}

func sum() {
	var num int
    fmt.Print("Enter the number : ")
	fmt.Scanln(&num)

	var sum int
	for i := 1; i <= num; i++ {
		sum += i
	}

	fmt.Printf("Sum of 1st %d is %d",num,sum)
}