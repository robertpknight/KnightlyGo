package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(speak("Hello", "World"))

	sentence, isName := speak("Hello", "W1rld")
	fmt.Println(sentence)
	fmt.Println(isName)
}

// func (r receiver) identifier(parameters) (return(s)) {...}
// when you call a function you pass arguments, when you define a function it is parameters

// speak ... take greeting and person, return a sentance and if the name was a name
func speak(greeting string, person string) (sentence string, isName bool) {
	sentence = fmt.Sprint(greeting, ", ", person)
	isName = isLetter(person)
	return sentence, isName
}

// check that all chars in string are alphabetic
func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
