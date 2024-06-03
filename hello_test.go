package hello_test

import (
	"fmt"
	"github.com/xiaoooooooooooo/hello"
	"testing"
)

func TestHello(t *testing.T) {
	data := ""
	except := fmt.Sprintf("hello %s!", data)

	result := hello.Hello(data)

	if result != except {
		t.Fatalf("except: %s, but got %s", except, result)
	}
}
