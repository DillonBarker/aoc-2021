package main

import (
	"aoc/commons/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Need to make a map m := make(map[int]int)
where x: y

x is 0-8, all the possible ages for the fish
and y is the number of fish at that age

		age: 0, count: 5
		age:1, count:6
		age:2, count:7
<<<< after a day passes >>>>
		age:0 , count: 6
		age:1, count: 7
		age:3, count: 0
		..
		age:6, count 5
		age:8, count 5


Then the sum of all these fish is the total fish
*/

func fishMap() map[int]int {
	m := make(map[int]int)
	fishes := getFishes()
	for _, fish := range fishes {
		m[fish] += 1
	}
	return m
}

func getTotalFishInBucket(bucket map[int]int) int {
	total := 0
	for i := 0; i <= 8; i++ {
		total += bucket[i]
	}
	return total
}

func main() {
	bucketOfFish := fishMap()
	for day := 0; day <= 255; day++ {
		temp8 := 0
		temp6 := 0
		for i := 0; i <= 8; i++ {

			if i == 0 {
				temp8 += bucketOfFish[0]
				temp6 += bucketOfFish[0]
				bucketOfFish[0] = 0
				continue
			}
			if i > 0 {
				bucketOfFish[i-1] += bucketOfFish[i]
				bucketOfFish[i] = 0
				continue
			}
		}
		bucketOfFish[8] += temp8
		bucketOfFish[6] += temp6
	}
	total := getTotalFishInBucket(bucketOfFish)
	fmt.Println(total)
}

func getFishes() []int {
	input, _ := os.Open("./day6/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	sc.Scan()
	inputs := strings.Split(sc.Text(), ",")
	fishes := []int{}
	for _, input := range inputs {
		inputAsInt, _ := helpers.ToInt(input)
		fishes = append(fishes, inputAsInt)
	}
	return fishes
}
