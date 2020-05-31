package main

import (
	"encoding/json"
	"fmt"
)

// weapon belt
type WeaponBelt struct {
	Knives int `json:"Knives"`
	Pistol int `json:"Pistol"`
}

type person struct {
	First      string `json:"First"`
	Last       string `json:"Last"`
	Age        int    `json:"Age"`
	WeaponBelt struct {
		Knives int `json:"Knives"`
		Pistol int `json:"Pistol"`
	} `json:"WeaponBelt"`
}

func main() {
	// produce person structs
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		WeaponBelt: WeaponBelt{
			Knives: 2,
			Pistol: 2,
		},
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		WeaponBelt: WeaponBelt{
			Pistol: 1,
		},
	}
	people := []person{p1, p2}
	fmt.Println("This is a people struct:\n", people)

	// marshal struct as JSON
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	s := string(bs)
	fmt.Println("\nthis is a JSON string:\n", s)

	// Unmarshall struct as JSON
	body := []byte(s)
	var agents []person

	err = json.Unmarshal(body, &agents)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nThis is the Unmarshalled JSON:\n", agents)

	// iterate over slice of type person
	fmt.Println("\nAccessing data from Unmarshalled JSON:")
	for i, agent := range agents {
		fmt.Println("Agent", i, ":")
		fmt.Println("\t", agent.First, agent.Last)
		fmt.Println("\t Carrying", agent.WeaponBelt.Knives, "Knives")
	}
}
