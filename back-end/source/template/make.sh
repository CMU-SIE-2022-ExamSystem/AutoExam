#!/bin/bash

# go to script's folder 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR

cd $1
# make autograde tar
tar cvf $SCRIPT_DIR/$1/autograde.tar autograder


cd ..
# make assessment's tar
tar cvf $SCRIPT_DIR/$1.tar $1