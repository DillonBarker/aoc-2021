#!/bin/bash

DAY=${1?Error: no day given eg. day1}

cd $DAY/part1 && go build main.go
cd ../..
cd $DAY/part2 && go build main.go

