package main

import (
	"fmt"
)

func Imprimir(a [2]string) {
	fmt.Print(a)
}

func RemoveElementIndex(array []int, value int) []int{
	b := append(array[:value], array[value+1:]...)
	return b
}

func ExerciseOne(dy, dx int) []int {
	a := make([]int, dy)
	for i := range a {
		a[i] = dx
	}
	return a
}

func ExerciseTwo(dy, dx int) [][]int {
	a := make([][]int, dy)
	for i := range a {
		a[i] = []int{dx}
	}
	return a
}

func ExerciseThree(dy, dx int) [][]int {
	a := make([][]int, dy)
	for i := range a {
		a[i] = make([]int, dx)
	}
	return a
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Hello"
	Imprimir(a)
	fmt.Println()

	slice := a[0:1]
	fmt.Print(slice)
	fmt.Println()

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	fmt.Println("r")
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	for i := 0; i < len(r); i++ {
		fmt.Print(r[i])
	}

	// A slice has both a length and a capacity.

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	fmt.Println("board", board)

	fmt.Println("\nRemoveElementIndex")
	array := []int{0,1,2,3,4,5,6,7,8,9}
	removed := RemoveElementIndex(array, 4)
	fmt.Printf("%v, %v\n\n", array, removed)


	res := ExerciseOne(3, 2)
	fmt.Println("ExerciseOne", res)					// ExerciseOne [2 2 2]

	res2 := ExerciseTwo(3, 2)
	fmt.Printf("ExerciseTwo %v - %T\n", res2, res2)	// ExerciseTwo [[2] [2] [2]] - [][]int

	res3 := ExerciseThree(3, 3)
	fmt.Printf("ExerciseThree %v - %T\n", res3, res3)	// ExerciseThree [[0 0 0] [0 0 0] [0 0 0]] - [][]int
}
