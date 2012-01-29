package clip

import (
	"os"
)

func init() {
	command["add"] = Add
}

func Add(args []string) (resp string, err os.Error) {
	if len(args) == 0 {
		err = os.NewError("nothing specified, nothing added")
		return
	}
	for _, arg := range args {
		library.AddPath(arg)
		resp += "Added " + arg + "\n"
	}
	return
}
