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
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		count += val
	}
	fmt.Println("Total: ", count)
}
