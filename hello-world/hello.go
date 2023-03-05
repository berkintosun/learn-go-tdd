package main

import "fmt"

const albanian = "alb"
const albanianPrefix = "Përshëndetje, "

const turkish = "tr"
const turkishPrefix = "Merhaba, "

const defaultPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case albanian:
		prefix = albanianPrefix
	case turkish:
		prefix = turkishPrefix
	default:
		prefix = defaultPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Beko", ""))
}
