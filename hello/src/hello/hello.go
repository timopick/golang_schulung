package main

import "fmt"

func main() {
	hello := "Hello World"

	fmt.Println(hello)
	fmt.Printf("slice: %v\n", hello[1:len(hello)-1])

	var olleh []byte
	for i := len(hello)-1; i >= 0; i-- {
		olleh = append(olleh, hello[i])
	}
	fmt.Println(olleh)

	execute("Zaphod", hellofunc)
}

func hellofunc(name string) {
	fmt.Println("Hello %v\n", name)
}

func execute(name string, f func(string)) {
	f(name)
}
