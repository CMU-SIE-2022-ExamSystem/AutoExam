#!/bin/bash

# install swag and initialize
FILE=~/go/bin/swag
if [ ! -f "$FILE" ]; then
    echo "install swag"  
   go install github.com/swaggo/swag/cmd/swag@latest
fi
echo "swag initialize"
~/go/bin/swag init

# # install air
FOLDER=./bin
if [ ! -d "$FOLDER" ]; then
    echo "install air"  
   curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
fi

# # run server
echo "run server"
go run main.go
# ./bin/air