package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i + 1
	}

	processNumber(arr)
}

func processNumber(arr []int) {
	for i := len(arr); i > 0; i-- {
		isPrime := isPrime(i)
		if (isPrime) {
			continue
		}

		toPrint := fooBar(i)
		fmt.Print(toPrint)

		if i != 1 {
			fmt.Print(", ")
		}
	}
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number % i == 0 {
			return false
		}
	}

	return true
}

func fooBar(number int) string {
	if number % 3 == 0 && number % 5 == 0 {
		return "FooBar"
	} else if number % 3 == 0 {
		return "Foo"
	} else if number % 5 == 0 {
		return "Bar"
	} else {
		return strconv.Itoa(number)
	}
}