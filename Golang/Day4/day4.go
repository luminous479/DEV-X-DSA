package day4
import "fmt"
//closure
func Outer() func() int {
	var count = 0
	return func() int {
		count++
		return count
	}
}
// struct
type Person struct {
	Name string
	Age  int
}
func PrintUserDetails(user Person) {
	fmt.Println("Name : ", user.Name)
	fmt.Println("Age : ",user.Age )

}
// receiver function
func(per Person) PrintPersonDetails(a string ) bool{
	fmt.Println(a)
	fmt.Println("Name : ", per.Name)
	fmt.Println("Age : ",per.Age )
	
	return true
}