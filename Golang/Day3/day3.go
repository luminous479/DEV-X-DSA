// higher order function
package day3

import "fmt"
func OperationProcess(a int, op func(c int)){ // higher order or call back function

	op(a)

} 

func PrintElement(x int){
  fmt.Println(" Higher Order function :",x)
  
}