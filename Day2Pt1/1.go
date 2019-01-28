package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inFile, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	twocount := 0
	threecount := 0
	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		str := scanner.Text()
		chars := make(map[rune]bool)
		twofound := false
		threefound := false

		for _, letter := range str {
			_, val := chars[letter]

			if !val {
				chars[letter] = true
				count := strings.Count(str, string(letter))
				if count == 2 {
					twofound = true
				}
				if count == 3 {
					threefound = true
				}
			}

		}

		if twofound {
			twocount++
		}
		if threefound {

			threecount++
		}
	}
	fmt.Println("Total: ", threecount*twocount)
}
