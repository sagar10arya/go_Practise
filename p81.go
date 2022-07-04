// Marshal to json

package main

import (
	"encoding/json"
	"fmt"
)

type user1 struct {
	First string
	Age   int
}

func main1() {
	u1 := user1{"James", 42}
	u2 := user1{"MoneyPenny", 27}

	people := []user1{u1, u2}
	fmt.Println(people)
	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
