package main

import (
	"fmt"
	"io/ioutil"
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

	_, err = f.Seek(0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	r, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", string(r))
}
