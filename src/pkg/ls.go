package clip

// This file implements the "ls" command.

import (
	"os"
	"bytes"
)

// Register the command.
func init() {
	command["ls"] = Ls
}

func Ls(args []string) (resp string, err os.Error) {
	buf := bytes.NewBuffer([]byte{})
	library.WriteTo(buf)
	resp = buf.String()
	return
}