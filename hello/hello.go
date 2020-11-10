package main

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return generateGreetingPrefix(language) + name
}

func generateGreetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	//implicitly returns variable from return type
	return
}

func main() {
	Hello("world", "english")
}
