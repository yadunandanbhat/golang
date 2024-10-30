package main

func Hello(name string) string {
	message := "Hello, "
	if name == "" {
		return message + "world"
	}
	return message + name
}