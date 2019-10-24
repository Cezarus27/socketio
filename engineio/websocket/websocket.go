package websocket

import (
	"github.com/3mdeb/socketio/engineio/transport"
)

var Creater = transport.Creater{
	Name:      "websocket",
	Upgrading: true,
	Server:    NewServer,
	Client:    NewClient,
}
