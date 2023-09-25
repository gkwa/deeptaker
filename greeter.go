package deeptaker

import (
	"fmt"
	"io"
	"log/slog"
)

func Greet(writer io.Writer, name string) {
	out := fmt.Sprintf("Hello, %s", name)
	slog.Debug(out)
	writer.Write([]byte(out))
}
