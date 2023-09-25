package deeptaker

import (
	"bytes"
	"fmt"
	"log/slog"
)

func Greet(writer *bytes.Buffer, name string) {
	out := fmt.Sprintf("Hello, %s", name)
	slog.Debug(out)
	writer.Write([]byte(out))
}
