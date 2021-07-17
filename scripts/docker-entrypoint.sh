#!/bin/bash

cmd=$1
case $cmd in
"bash")
    bash
    ;;
"test")
    CompileDaemon -build="go version" -command="go test -coverprofile=cover.out ./..."
    ;;
"build" | *)
    CompileDaemon -build="go build -o bin/user main.go" -command="bin/user"
    ;;
esac
