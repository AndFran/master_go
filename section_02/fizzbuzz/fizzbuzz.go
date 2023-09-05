package main

import (
	"fmt"
	"os"
)

func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case n%15 == 0:
			fmt.Println("FizzBuzz")
		case n%3 == 0:
			fmt.Println("Fizz")
		case n%5 == 0:
			fmt.Println("Buzz")
		}
	}
}

func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := fmt.Scanf("%d", &os.Args[1])
		if err == nil {
			n = max
		} else {
			fmt.Println(err)
			return
		}
	}
	fizzbuzz(n)
}
