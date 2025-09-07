package main

import "fmt"



func getCompleteName()(firstName, midleName, lastName string){
	firstName = "Moch"
	midleName = "Abdul"
	lastName = "Azis"

	return firstName, midleName, lastName
}

func main(){
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
