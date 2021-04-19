package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type ByAge []Person
type ByFirstName []Person
type ByLastName []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByFirstName) Len() int           { return len(a) }
func (a ByFirstName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirstName) Less(i, j int) bool { return a[i].FirstName < a[j].FirstName }

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastName) Less(i, j int) bool { return a[i].LastName < a[j].LastName }

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println("Unsorted:", people)

	// sort by age
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("By age:", people)

	// sort by first name
	sort.Sort(ByFirstName(people))
	fmt.Println("By first name:", people)
}
