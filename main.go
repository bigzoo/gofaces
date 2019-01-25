package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct {}
type spanishBot struct {}

func (englishBot) getGreeting() string{
	// very very custom ish logic for returning greeting
	return "Hello world"
}

func (spanishBot) getGreeting() string {
	//very very custom login for spanish
	return "Hola"
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}


func main (){
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}