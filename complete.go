package main

// This file implements bash programmable completion.

import (
	"fmt"
	"strconv"
)

func init() {
	help["complete"] = `do bash autocompletion on arguments`
}

// args: $COMP_WORDC $COMP_LINE, e.g.:
//	3 clip play jaz
func (api API) Complete(args []string) (resp, err string) {
	// don't crash on panic
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Sprint(e)
			resp = ""
		}
	}()

	// arg[0]: $COMP_WORDC: index of word to complete
	// arg[i]: $COMP_LINE[i-1]: "clip arg0 arg1 ..."
	myargs := []string{}
	if len(args) > 2 {
		idx, _ := strconv.Atoi(args[0])
		myargs = args[2 : idx+2]
	}
	Debug("complete", myargs)

	return
}

// Used by "clip -c", invoked by bash completion.
// args:
//	word line
// E.g.:
//	clip alpha beta<TAB>
// yields args:
//	beta clip alpha beta
//func AutoComplete(args []string) {
//	if len(args) == 0 {
//		return // should not happen
//	}
//	if len(args) == 1 {
//		// fix for word = "" (omitted by bash)
//		args = []string{"", args[0], ""}
//	}
//	word := args[0]
//	//cmd := args[1]
//	//line := args[2:]
//	if len(args) == 3 {
//		completeCommands(word)
//		return
//	}
//}
//
//
//// Auto-complete function for player commands like 
////	add ls play ...
//func completeCommands(prefix string) {
//	for cmd, _ := range player.command {
//		if strings.HasPrefix(cmd, prefix) {
//			fmt.Print(cmd, " ")
//		}
//	}
//}
