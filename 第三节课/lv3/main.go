package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("plan.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = f.Write([]byte("I’m not afraid of difficulties and insist on learning programming"))
	if err != nil {
		log.Fatalln(err)
	}
}
