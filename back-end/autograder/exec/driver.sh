#!/bin/bash

# Usage: ./driver.sh

# Run the autograder
echo "<---Running--->"
py=`python3.8 main.py $1`
status=$?
if [ ${status} -eq -1 ]; then
    echo "<---Failure--->"
else
    echo "<---Success--->"
fi
echo $py

exit
