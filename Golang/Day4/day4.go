package day4

func Outer() func() int{

	var count = 0;

	return func() int{
		count++
		return count
	}
}