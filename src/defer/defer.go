// defer call calculate all the arguments of a method right away
// but doesnt call the method until the wrapped method returns

package main

import "fmt"

func doSome() string {
	fmt.Println("doing something")
	return "111"
}

func increment() (i int) {
    defer func() { i++ }()
    return 1
}

func main() {
	defer fmt.Println("world", doSome())
	fmt.Println("hello")

	fmt.Println("Defer can change the value of a named return variable: ", increment())
}
