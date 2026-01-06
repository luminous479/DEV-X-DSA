package main

import "fmt"

func add(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println(sum)
}

func main() {

	fmt.Println("hello world!")

	// variable declaration 
	var a int = 10
    fmt.Println(a)
	 b := 5
	fmt.Println(b)
	b = 15
	fmt.Println(b)
	const p = 100;
    fmt.Println(p)

	// if....else and switch
    
	age := 10
	sex := "female"

	if age >= 18 && sex == "male"{
		fmt.Println("university male Student")
	} else if age < 18 && sex == "female"{
		fmt.Println("Hi school female Student")
	} else {
		fmt.Println("kintergarden Student")
	}

	num := 5

	switch num{
	case 1: {
		fmt.Println("one")
	}	
	case 2: {
		fmt.Println("two")
	}
	default : {
		fmt.Println("more then 2")
	}	
	}

	// functions
	x := 10
	y := 20

	add(x, y)

}