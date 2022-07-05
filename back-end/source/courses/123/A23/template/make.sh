#!/bin/bash

# go to script's folder 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR

# make assessment's tar
tar cvf $SCRIPT_DIR/$1.tar $1