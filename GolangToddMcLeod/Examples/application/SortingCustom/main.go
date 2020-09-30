package main

import (
	"fmt"
	"sort"
)

// sorting structs will require implementing the Interface interface in the Sort package

type person struct {
	First string
	Age   int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByFirst implements sort.Interface for []Person based on
// the First field.
type ByFirst []person

func (f ByFirst) Len() int           { return len(f) }
func (f ByFirst) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByFirst) Less(i, j int) bool { return f[i].First < f[j].First }

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println("Sort by First Name:")
	fmt.Println(people)
	sort.Sort(ByFirst(people))
	fmt.Println(people)

	fmt.Println("\nSort by Age:")
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
