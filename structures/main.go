package main

import (
	"fmt"
	"reflect"
)

type Sponsor struct{
	name string
}

type Equipo struct{
	name string `required max="25"`
	sponsor Sponsor
}

type Torneo struct{
	equipos []Equipo
	sponsor Sponsor
}

func main(){
	fmt.Println("Structures")
	fmt.Printf("")

	bbva := Sponsor{
		"BBVA",
	}
	
	river := Equipo{
		"River",
		bbva,
	}

	boca := Equipo{
		"Boca",
		bbva,
	}

	torneo := Torneo{
		equipos: []Equipo{river, boca},
		sponsor: bbva,
	}

	fmt.Printf("Torneo: %v",torneo)
	fmt.Printf("Sponsor del torneo: %v",torneo.sponsor)
	fmt.Println("")

	t := reflect.TypeOf(Equipo{})
	field, _ := t.FieldByName("name")
	fmt.Printf("Tag: %v",field.Tag)

	fmt.Println("")

}