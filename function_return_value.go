package main

import "fmt"

func getFullName() (string,string){
	return "Abdul", "Azis"
}

func main(){
	firstName, _ := getFullName()
	fmt.Println(firstName)
}