package main

import (
	"fmt"
	"time"
)

func task(i int) {

	fmt.Println(i)
}

func main() {

	for i := 1; i <= 10; i++ {
		go task(i)
	}

	time.Sleep(2 * time.Second)

}