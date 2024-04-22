package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	fmt.Print("Loop: ")
	var loopInput string
	fmt.Scan(&loopInput)
	//
	//fmt.Print("Threads: ")
	//var threadInput string
	//fmt.Scan(&threadInput)

	start := time.Now()

	initCalc(stringToInt(loopInput))

	elapsed := time.Since(start)
	fmt.Printf("Benchmark Finished after: %v", elapsed)
}

func initCalc(loop int) {
	//var multiplier = 0.5
	//var tempStartNumber = startNumber
	//var tempEndNumber = endNumber
	//
	//for i := 0; i < threads; i++ {
	//	var finalStartNumber int
	//	var finalEndNumber int
	//
	//	finalStartNumber = tempStartNumber
	//	finalEndNumber = int(math.Round((endNumber / i) * multiplier))
	//}
	calc(0, loop)
}

func calc(startNumber int, endNumber int) {
	fmt.Println("Started...")

	for i := startNumber; i < endNumber; i++ {
		var number1 = getDivisorSum(i)
		var number2 = getDivisorSum(number1)

		if number2 == i && number1 != number2 {
			fmt.Printf("Result: %v - %v \n", number1, number2)
		}
	}
}

func getDivisorSum(value int) int {
	divisorSum := 0
	for i := 1; i <= int(math.Sqrt(float64(value))); i++ {
		if value%i == 0 {
			divisorSum += i
			if i != value/i {
				divisorSum += value / i
			}
		}
	}
	return divisorSum - value
}

func stringToInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		// ... handle error
		panic(err)
	}

	return i
}
