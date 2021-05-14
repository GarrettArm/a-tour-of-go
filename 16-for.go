package main

import "fmt"

func main() {
	sum := 1
	/*
		for init; condition; post {
			enclosed scope
		}
		for i := 0; i < 10; i++ {
			sum += i
		}
	*/
	/*
		init & post are optional
		becomes 'while' loop
	*/
	/*
		for sum < 1000 {
			sum += sum
		}
	*/
	/*
		init, condition, post are optional
		becomes infite loop
	for {
		fmt.Println("inside loop")
	}
	fmt.Println(sum)
}
