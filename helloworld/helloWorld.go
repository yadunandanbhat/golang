package main

const (
	spanish = "spanish"
	french  = "french"

	spanishPrefix = "Hola, "
	englishPrefix = "Hello, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	message := getPrefix(language)
	return message + appendMessage(name)
}

func getPrefix(language string) string {
	switch language {
	case spanish:
		return spanishPrefix
	case french:
		return frenchPrefix
	default:
		return englishPrefix
	}
}

func appendMessage(name string) string {
	if name == "" {
		return "world"
	}
	return name
}