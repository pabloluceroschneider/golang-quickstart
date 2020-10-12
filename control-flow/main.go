package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	// defer LIFO

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close() 
	// Close() should be at the end of the program
	// but this keyborad allows to have the open and close next to each other
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

}