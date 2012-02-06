package clip

// This file implements the Remote Procedure Call between
// the clip daemon and front-ends

import (
	"os"
	"rpc"
	"net"
	"http"
	"reflect"
	"strings"
)

// Start serving RPC calls from client instances.
func (player *Player) serveRPC() {
	rpc.Register((*PlayerRPC)(player))
	rpc.HandleHTTP()
	conn, err := net.Listen("tcp", port)
	if err != nil {
		Err("listen error:", err)
	}
	Debug("Listening on port " + port)
	http.Serve(conn, nil)
	//TODO: log errors.
}

// Dummy type to define RPC methods on.
type PlayerRPC Player


// RPC-exported function used for normal operation mode.
// The command-line arguments are passed (e.g. "play jazz")
// and a response to the user is returned in *resp.
func (rpc *PlayerRPC) Call(args []string, resp *string) (err os.Error) {
	Debug("PlayerRPC.Call", args)

	player := (*Player)(rpc)

	// utf8.NewString(...)
	cmd := unicode.ToUpper(args[0]) + args[0][1:]
	args = args[1:]

	p := reflect.ValueOf(player)
	m := p.MethodByName(cmd)
	Debug("MethodByName", cmd, ":", m)
	if m.Kind() == reflect.Invalid{
		err = os.NewError("clip: '" + cmd + "' is not a clip command. See clip help.")
		return
	}
	r := m.Call([]reflect.Value{reflect.ValueOf(args)})
	*resp = r[0].Interface().(string)   // by convention, response is 1st return value
	errStr := r[1].Interface().(string) // by convention, error is 2nd return value
	if errStr != ""{err = os.NewError(errStr)}

	return
}
