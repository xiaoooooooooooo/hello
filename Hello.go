package hello

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "go"
	}
	return fmt.Sprintf("hello %s!", name)
}
