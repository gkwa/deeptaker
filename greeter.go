package deeptaker

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	out := []byte(fmt.Sprintf("Hello, %s", name))
	writer.Write(out)
}
