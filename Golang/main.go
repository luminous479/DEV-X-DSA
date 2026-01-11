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
   
   // day 4 : closure , struct, receiver function 
      incre1 := day4.Outer()
	  fmt.Println(incre1())
	  fmt.Println(incre1())

     person1 := day4.Person{
		Name : "Jhon",
		Age :  30,
	 }
	 user1 :=  day4.Person{
	Name : "Tina",
	Age : 20,
   }
   user2 :=  day4.Person{
	Name : "Hina",
	Age : 21,
   }
	 fmt.Println(person1)

	 day4.PrintUserDetails(user1)
	 // calling receiver function
	 isPrinted := user2.PrintPersonDetails("=====Person Details=======")
	 fmt.Println(isPrinted)

	 // array 
	 arr := [5]int{1, 2, 3, 4, 5} // Short declaration with initial values
	 fmt.Println(arr)
	var arr1 [3]int
	 arr1[0] = 3
	 arr1[1] = 5
	 fmt.Println(arr1)
	 var twoD [2][2]int

	 twoD = [2][2]int{
		{1,2},
		{3,4},
	 }
	 fmt.Println(twoD)
	 //pointer
	 x := 20
	 fmt.Println(x)
	 p := &x
	 *p = 30
	 fmt.Println(*p)

	  num := [3]int{1,2,3}
 
	  printnum(&num)


}
func printnum(num *[3]int){
        fmt.Print(num)
	  }

func init(){
	fmt.Println("this is init value",i)
	i = 20
}

