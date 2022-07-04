// Encode to json the []user sending the results to stdout.
// Hint: You will need to use json.NewEncoder(os.Stdout).encode(v interface{})

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {

	u1 := user{First: "James", Last: "Bond", Age: 32, Sayings: []string{"Shaken, not stirred",
		"Youth is no guarantee for innovation", "In his majesty's royal service"}}
	u2 := user{First: "Miss", Last: "MoneyPenny", Age: 27, Sayings: []string{"James, its so good to see you",
		"Would you like me to take care of that for you,James",
		"I would really prefer to be secret agent myself,"}}
	u3 := user{First: "M", Last: "Hmmm", Age: 54, Sayings: []string{"Oh James you didn't",
		"Dear god, what has james done now",
		"Can someone tell me where james bond is?,"}}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We have done something wrong", err)
	}

}
