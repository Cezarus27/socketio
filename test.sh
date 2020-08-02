#!/bin/bash

go test -race

cd engineio
go test -race

cd polling
go test -race
cd -

cd websocket
go test -race
cd -

cd parser
go test -race
cd -
