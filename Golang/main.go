package main

import (
	"fmt"

	day2 "github.com/luminous479/DEV-X-DSA/Day2"
	day3 "github.com/luminous479/DEV-X-DSA/Day3"
	day4 "github.com/luminous479/DEV-X-DSA/Day4"
)
var (
	a = 10 // global
	b = 20 // global
)
var p = 10
var i = 5
func main() {
	day2.Add(a,b)
    var age = 30
	// variable shadowing
	if (age>5){
		p := 20
		fmt.Println("NOT KID",p)
		
	}
	fmt.Println(p)
	fmt.Println("this is main value",i)

	// anonymous function and IIFE
	func(a string){
		fmt.Println("This is ", a ," function ")
	}("anonymous")

	// a noob function with expression

	txt := func ( n string){ // parameter 
       
		fmt.Println("THis is a noob ",n) 
	}
	txt("function") // argument
   day3.OperationProcess(30,day3.PrintElement)
   //
      incre1 := day4.Outer()
	  fmt.Println(incre1)
	  fmt.Println(incre1)

	
}

func init(){
	fmt.Println("this is init value",i)
	i = 20
}

