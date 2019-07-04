// This code is an example of defer-panic-recover functions
// when the input value is greater than 5, the validate function panics! 
// panic reaches to the startBusinessFlow() function and it in tern calls its deferred function
// defer function then recovers and finds what went wrong and returns nice readable output to the caller

package main

import "fmt"

func main () {
	fmt.Println("Demo of defer-panic-recover!")
	
	var output string 
	fmt.Println("\n *** POSITIVE FLOW: ***")
	output = startBusinessFlow(3)
	fmt.Printf("Output from business flow = %v \n", output)

	fmt.Println("\n *** NEGATIVE FLOW: ***")
	output = startBusinessFlow(7)
	fmt.Printf("Output from business flow = %v \n", output)

}

func startBusinessFlow(input int) (output string) {
	fmt.Printf("Input to the business flow = %v \n", input)

	defer func () {
		//this function handlers panic errors
		r := recover()
		if r != nil {
			fmt.Printf("Something went wrong. Here's the error message: %q \n", r)
			output = "INVALID"
		}
	}()	

	output = performOperation(input)
	return
}

func performOperation(input int) (output string) {
	fmt.Println("Starting operation...")
	
	validate(input)

	fmt.Printf("Operation is to print square of %v = %v \n", input, input*input)
	return "SUCCESS"
}

func validate(input int) {
	if input > 5 {
		panic("Invalid input value greater than 5")
	}
}