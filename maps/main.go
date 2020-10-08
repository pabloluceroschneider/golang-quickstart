package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hola")

	nombres := map[int]string{
		1: "Pablo",
		2: "Daniel",
		3: "Lucero",
		4: "Schneider",
	}
	fmt.Println("Primer Nombre: ", nombres[1])  	//Pablo
	fmt.Println("Nombres: ", nombres)				//Pablo Daniel Lucero Schneider

	// delete(nombres, 2)
	// fmt.Println("Nombres: ", nombres)			//Pablo Lucero Schneider

	_, ok := nombres[5]								// _ devuelve 0; ok devuelve booleano
	fmt.Println("¿Tiene quinto nombre?:", ok)		// Quinto nombre


	copy := nombres
	delete(nombres, 2)
	fmt.Println("Copy: ", copy, nombres)			//Pablo Lucero Schneider (modifica ambos)



}