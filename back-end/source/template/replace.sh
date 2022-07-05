#!/bin/bash

# go to script's folder 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR

# change folder name
mv temp $1

cd $1

# replace name
mv temp.rb $1.rb
mv temp.yml $1.yml

# replace .rb content
sed -i'' -e "s/Temp/$1/g" $1.rb
sed -i'' -e "s/temp/$1/g" $1.rb

# make autograde tar
tar cvf $SCRIPT_DIR/$1/autograde.tar autograder
