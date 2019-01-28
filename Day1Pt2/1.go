package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inFile, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	count := 0
	countMap := make(map[int]bool)
	var values []int
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		values = append(values, val)
	}
	for {
		for _, val := range values {
			countMap[count] = true
			count += val
			//fmt.Println(count)
			if countMap[count] {

				fmt.Println("First repeat: ", count)
				os.Exit(0)
			}
		}

	}
}
