package socketio

import (
	"os"
	"strconv"
)

const (
	db1 = false
)

var (
	DbLogMessage = false
	LogMessage, _ = strconv.ParseBool(os.Getenv("SOCKETIO_LOG_MESSAGE"))
)
