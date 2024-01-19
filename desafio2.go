package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPam")
		} else if i%3 == 0 {
			fmt.Println("Pin")

		} else if i%5 == 0 {
			fmt.Println("Pin")
		} else {
			fmt.Println(i)
		}
	}

}
