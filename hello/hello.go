package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishGreetingsPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishGreetingPrefix
	case french:
		return frenchGreetingPrefix
	default:
		return englishGreetingsPrefix
	}
}

func main() {
	fmt.Println(Hello("", ""))
}
