package socketio

import (
	"os"
	"strconv"
)

const (
	db1 = false
)

var (
	DbLogMessage, _ = strconv.ParseBool(os.Getenv("SOCKETIO_LOG_MESSAGE"))
	LogMessage, _ = strconv.ParseBool(os.Getenv("SOCKETIO_LOG_MESSAGE"))
)
