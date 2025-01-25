package di

import (
	"fmt"
	"bytes"
)

func Greeter(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}