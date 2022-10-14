/*
	It shows an example of how to sort slices with package "sort"
	in the standard library (https://pkg.go.dev/sort):

		Package sort provides primitives fork
		sorting slices and user-defined collections.

	The user-defined type shall implement sort.Interface
	to make use of the functions in this package.
	type Interface interface {
		Len() int
		Less(i, j int) bool
		Swap(i, j int)
	}

	Important functions:
		func Sort(data Interface)
		func Stable(data Interface)

	Sort sorts data in ascending order as determined by
	the Less method. It makes one call to data.Len to
	determine n and O(n*log(n)) calls to data.Less and data.Swap.
	The sort is not guaranteed to be stable.
	The Stable version keeps the original order of equal elements.

*/
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

}

/*
Output:

[Bob: 31 John: 42 Michael: 17 Jenny: 26]
[Michael: 17 Jenny: 26 Bob: 31 John: 42]
[John: 42 Bob: 31 Jenny: 26 Michael: 17]

*/
