//From sample
package sample
package mainn
package main

import "fmt"

var name string
var designation string

func takeName() {

	fmt.Printf("Enter your name :")
	fmt.Scanln(&name)
	fmt.Printf("The name entered is : %s", name)
}

func takeDesignation() {

	fmt.Printf("\nEnter your designation :")
	fmt.Scanln(&designation)
	fmt.Printf("Your designation is : %s", designation)
}

func calcAvgHeight() {
	var n int
	var avg float32 = 0
	var height float32
	fmt.Println("\nEnter n for calculating height:")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&height)
		avg = avg + height
	}

	avg = (avg) / float32(n)

	fmt.Printf("The avg height is %f", avg)

}

func createArray() {
	var n int
	fmt.Println("\nEnter n for creating array:")
	fmt.Scanln(&n)
	var arr1 [10]int
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr1[i])
	}
	fmt.Println("\nThe entered numbers are :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr1[i])
	}
}

func createStringArray() {
	var n int
	fmt.Println("\nEnter n for creating array of strings:")
	fmt.Scanln(&n)
	var arr1 [10]string
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr1[i])
	}
	fmt.Println("\nThe entered Strings are :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s ..", arr1[i])
	}
}

func printUsingFor() {
	var n int
	fmt.Println("\nEnter n for printing upto n:")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i+1)
	}
}

func sumUsingRange() {
	var n int
	fmt.Println("\nEnter n for sum upto n:")
	fmt.Scanln(&n)
	var arr1 [10]int
	for i := 0; i < n; i++ {
		arr1[i] = i + 1
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d ", arr1[i])
	// }

	var sum int = 0

	for k := range arr1 {
		sum = sum + arr1[k]
	}
	fmt.Printf("The sum upto %d is %d", n, sum)
}

func findNum() {
	var flag bool = true

	var n int = 9
	for flag {
		if n%2 == 1 && n%3 == 1 && n%4 == 1 && n%5 == 1 && n%6 == 1 && n%7 == 0 {
			fmt.Printf("\nThe number is %d", n)
			//time.sleep(0.1)
			// flag = false

		}
		n++

	}

}

func guessNumber() {
	fmt.Println("Guess number from 1 to 10: ")
	var n int
	fmt.Scanln(&n)
	var count int = 1
	var secretNumber int = 5
	for count < 3 && secretNumber != n {
		if n < secretNumber {
			fmt.Println("You are wrong...Guess a higher number ")
		} else {
			fmt.Println("You are wrong...Guess a lower number ")

		}
		fmt.Scanln(&n)
		count++
	}
	if count <= 3 && secretNumber == n {
		fmt.Printf("You won at the %d attempt/attempts", count)
	} else {
		fmt.Println("You lost dude")
	}

}

func main() {
	// takeName()

	// //This is a comment

	// takeDesignation()

	// calcAvgHeight()

	// createStringArray()
	// printUsingFor()
	// sumUsingRange()

	//findNum()
	guessNumber()
}
