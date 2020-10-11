package main

import "fmt"

func main(){
	// elemental for
	for i := 0 ; i < 5; i++{
		fmt.Println("i:",i)
	}

	// for two variables
	for i,j := 0,0 ; i < 5; i, j = i+1, j+2{
		fmt.Println("i,j:",i,j)
	}

	// sugar syntax
	a := 0
	for {
		fmt.Println(a)
		a++
		if a == 5 {
			break
		}
	}

	// Loop label
	Loop: 
		for i:=0;i<5;i++{
			for j:=0;j<5;j++{
				fmt.Println("i*j:", i*j)
				if i*j >= 3 {
					break Loop
				}
			}
		}
	
	// Collections
	s := []int{1,2,3}
	// k,v is key and value
	for k, v := range s {
		fmt.Println(k,v)
	}
}