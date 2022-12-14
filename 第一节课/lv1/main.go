package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func readLine(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	return reader.ReadString('\n')
}

func main() {
	input, err := readLine(os.Stdin)
	if err != nil {
		log.Fatalf("failed to readLine: %v\n", err)
	}
	input = input[0 : len(input)-1]

	rInput := []rune(input)
	rCount := len(rInput)
	for i := 0; i < rCount; i++ {
		if rInput[i] != rInput[rCount-1-i] {
			println("不是回文")
			os.Exit(0)
		}
	}
	if len(input)%2 != 0 {
		fmt.Printf("是回文: %s\n", string(rInput[:len(rInput)/2+1]))
	} else {
		fmt.Printf("是回文: %s\n", string(rInput[:len(rInput)/2]))
	}
}
