package main

import "fmt"

// go mod init something

const (
	spanish = "Spanish"
	french = "French"
	ukraine = "Ukraine"


	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	ukraineHelloPrefix = "Pruvit, "
)



func Hello(name string, language string) string {	
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix (language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case ukraine:
		prefix = ukraineHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}