package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		users = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
		distribution = make(map[string]int)
	)

	total := 0
	for _, user := range users {
		lowerUser := strings.ToLower(user)
		coins := 0
		coins += strings.Count(lowerUser, "e")
		coins += strings.Count(lowerUser, "i") * 2
		coins += strings.Count(lowerUser, "o") * 3
		coins += strings.Count(lowerUser, "u") * 4
		distribution[user] = coins
		total += coins
	}

	for k, v := range distribution {
		fmt.Printf("%v获得%v个金币\n", k, v)
	}

	fmt.Printf("一共花费%v个金币\n", total)
}
