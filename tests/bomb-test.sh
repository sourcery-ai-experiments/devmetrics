#!/bin/bash

for ROW in {1..10000}
do
    echo "Iteration: $ROW"
    PollCount="$(curl -s -X GET http://localhost:8080/value/counter/PollCount)"
    echo "value: PollCount = $PollCount"
    echo "###########################"
    # printf '\n'
done
