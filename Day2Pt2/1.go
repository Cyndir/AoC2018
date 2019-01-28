package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inFile, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	var boxes []string

	for scanner.Scan() {
		str := scanner.Text()
		boxes = append(boxes, str)
	}
	for i, box := range boxes {
		for diffIndex := range box {
			match := box[:diffIndex] + box[diffIndex+1:]

			for _, nextBox := range boxes[i+1:] {
				if match == nextBox[:diffIndex]+nextBox[diffIndex+1:] {
					fmt.Println("ID:", match)
					fmt.Println("Box 1", box)
					fmt.Println("Box 2", nextBox)
					os.Exit(0)
				}
			}
		}

	}

}
