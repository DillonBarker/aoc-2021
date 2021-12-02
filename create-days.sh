#!/bin/bash

for DAYNUMBER in {1..25} 
do
	mkdir "day${DAYNUMBER}"
	cd "day${DAYNUMBER}" && mkdir "part1" && mkdir "part2" && touch "input.txt"
  cd ..
done
