package day1

import "fmt"

func add(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println(sum)
}
func calculator(num1 int, num2 int) (int, int){

	sum := num1 + num2
	mul := num1 * num2

	return sum, mul

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
   // function with return values and types

   s,m := calculator(5,7)

   fmt.Println(s)
   fmt.Println(m)
}