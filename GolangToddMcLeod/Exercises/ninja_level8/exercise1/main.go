package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println("Print users struct:", users)

	// your code goes here

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nPrint slice of bytes:", bs)

	// unmarshall the slice of bytes to unmarshalUsers
	var unmarshalUsers []user

	err = json.Unmarshal(bs, &unmarshalUsers) // unmarshall requires a point to make changes to the variable in place
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nThis is the unmarshalled slice of user:", unmarshalUsers)

}
