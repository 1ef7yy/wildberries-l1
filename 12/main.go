package main

import "fmt"

type StringSet struct {
	data map[string]bool
}

func NewStringSet() *StringSet {
	return &StringSet{data: make(map[string]bool)}
}

func (s *StringSet) Add(value string) {
	s.data[value] = true
}

func (s *StringSet) Print() {
	fmt.Println("Set elements:")
	for value := range s.data {
		fmt.Println(value)
	}
}

func CreateSetFromSlice(slice []string) *StringSet {
	set := NewStringSet()
	for _, value := range slice {
		set.Add(value)
	}
	return set
}

func main() {
	set := CreateSetFromSlice([]string{"cat", "cat", "dog", "cat", "tree"})
	set.Print()
}
