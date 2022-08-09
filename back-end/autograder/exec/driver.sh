#!/bin/bash

# Usage: ./driver.sh

# Run the autograder
<<<<<<< HEAD
source env/bin/activate
echo "<---Running--->"
py=`python3.9 main.py $1`
=======
echo "<---Running--->"
py=`python main.py $1`
>>>>>>> 6a1b6dc (Add functions for process modules in customized autograders.)
status=$?
if [ ${status} -eq -1 ]; then
    echo "<---Failure--->"
else
    echo "<---Success--->"
fi
echo $py

exit
<<<<<<< HEAD

=======
>>>>>>> 6a1b6dc (Add functions for process modules in customized autograders.)
