package main

import "fmt"

// неупорядоченность элементов дает нам реализовать множество через map
type Set struct {
	data map[interface{}]bool
}

func NewSet() *Set {
	return &Set{data: make(map[interface{}]bool)}
}

// при добавлении элемента в множество устанавливаем его значение в true,
// что позволяет сохранить уникальность значений
func (s *Set) Add(value interface{}) {
	s.data[value] = true
}

func (s *Set) Print() {
	fmt.Println("Set elements:")
	for value := range s.data {
		fmt.Println(value)
	}
}

func FindIntersection(s1, s2 *Set) *Set {
	intersection := NewSet()
	// проходимся по первому множеству
	for value := range s1.data {
		// проверяем есть ли элемент во втором множестве
		if _, ok := s2.data[value]; ok {
			intersection.Add(value)
		}
	}
	return intersection
}

func main() {
	s1 := NewSet()
	s2 := NewSet()
	for i := 1; i <= 10; i++ {
		s1.Add(i)
	}

	for i := 4; i <= 8; i++ {
		s2.Add(i)
	}

	intersection := FindIntersection(s1, s2)
	intersection.Print()
}
