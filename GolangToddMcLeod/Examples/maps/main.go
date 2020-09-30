package main

import (
	"fmt"
)

func main() {
	// a map is also a composite literal
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])

	// keys that do not exist will still return zero
	fmt.Println(m["Barnabas"])

	// comma ok idiom
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	// add elements to maps
	fmt.Println("adding rob to map")
	m["Rob"] = 27
	fmt.Println(m)

	// delete elements from maps
	fmt.Println("removing rob from map")
	delete(m, "Rob")
	fmt.Println(m)

	// removing a key that doesn't exist will not error
	delete(m, "Todd")

	// adding to conditional logic
	if v, ok := m["James"]; ok {
		fmt.Println("this checks that James exists:", v)
	}

	// check that what you want to delete exists
	if _, ok := m["Todd"]; ok {
		fmt.Println("trying to delete")
		delete(m, "Todd")
	}

	// iterate over map using range
	for k, v := range m {
		fmt.Println(k, v)
	}
}
