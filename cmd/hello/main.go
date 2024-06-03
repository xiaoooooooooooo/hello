package main

import (
	"flag"
	"github.com/xiaoooooooooooo/hello"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "go", "命令行打印hello")
}
func main() {
	flag.Parse()
	msg := hello.Hello(name)
	_, err := os.Stdout.WriteString(msg)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
