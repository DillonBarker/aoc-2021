#!/bin/bash

for DAYNUMBER in {1..25} 
do
  if [[ -f "./day$DAYNUMBER/part1/main" ]] && [[ -f "./day$DAYNUMBER/part2/main" ]]
  then
    echo "$DAYNUMBER:
      part1: $(./day$DAYNUMBER/part1/main)
      part2: $(./day$DAYNUMBER/part2/main)
    "
  fi
done