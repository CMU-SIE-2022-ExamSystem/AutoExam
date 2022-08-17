#!/bin/bash

# Usage: ./driver.sh

# Run the autograder
echo "<---Running--->"
python3.9 main.py
status=$?
if [ ${status} -eq -1 ]; then
    echo "<---Failure--->"
# else
#     echo "<---Success--->"
fi

exit
