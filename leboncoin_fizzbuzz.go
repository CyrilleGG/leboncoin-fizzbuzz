package main

import (
	"fmt"
	"strconv"
)

var count int = 1
var limit int
var int1 int
var int2 int
var str1 string
var str2 string
var results []string

func main() {
	limit = 30
	int1 = 3
	int2 = 5
	str1 = "fizz"
	str2 = "buzz"

	for count <= limit {
		if count%int1 == 0 && count%int2 == 0 {
			results = append(results, str1+str2)
		} else if count%int1 == 0 {
			results = append(results, str1)
		} else if count%int2 == 0 {
			results = append(results, str2)
		} else {
			results = append(results, strconv.Itoa(count))
		}

		count++
	}

	fmt.Println(results)
}