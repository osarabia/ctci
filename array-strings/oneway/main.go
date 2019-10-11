package main

import "fmt"

/*
There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character.Given two strings, write a function to check if the are one edit
 (or zero edits) away
*/

func main() {
    s1 := []string{"p", "a", "l", "e"}
    s2 := []string{"p", "l", "e"}

    r := oneWay(s1, s2)
    fmt.Printf("%v\n", r)

    s1 = []string{"p", "a", "l", "e", "s"}
    s2 = []string{"p", "a", "l", "e"}

    r = oneWay(s1, s2)
    fmt.Printf("%v\n", r)

    s1 = []string{"p", "a", "l", "e"}
    s2 = []string{"b", "a", "l", "e"}

    r = oneWay(s1, s2)
    fmt.Printf("%v\n", r)

    s1 = []string{"p", "a", "l", "e"}
    s2 = []string{"b", "a", "k", "e"}

    r = oneWay(s1, s2)
    fmt.Printf("%v\n", r)

    s1 = []string{"p", "a", "l", "e", "e"}
    s2 = []string{"p", "a", "l", "e"}

    r = oneWay(s1, s2)
    fmt.Printf("%v\n", r)

    s1 = []string{"b", "a", "l", "e", "e"}
    s2 = []string{"p", "a", "l", "e"}

    r = oneWay(s1, s2)
    fmt.Printf("%v\n", r)


}

func oneWay(s1 []string, s2 []string) bool {
	s2Counter := make(map[string]int)

	changes := 0

	for _,v := range s2{
		s2Counter[v] = s2Counter[v] + 1
	}


	for _,v1 := range s1 {
		_,ok := s2Counter[v1]
		if ok{
			if s2Counter[v1] > 0 {
				s2Counter[v1] = s2Counter[v1] - 1
			} else {
				if changes > 0 {
					return false
				}
				changes += 1
			}
		} else {
			if changes > 0 {
				return false
			}
			changes += 1
		}
	}

	return true
}
