package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	message := "Hello Pablo"
	i, j := 5, 10

	var p *int
	p = &i

	d := &j

	*d = 21

	person := User{"Pablo", 24}

	fmt.Println(*p)
	fmt.Println(*d)
	fmt.Println(message)
	fmt.Println(j)
	fmt.Println(person.Name)
}
