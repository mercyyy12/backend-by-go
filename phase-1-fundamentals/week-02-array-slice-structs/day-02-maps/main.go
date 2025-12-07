package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		// empty map
		// m := map[string]int{} //map declaration
		// m["Mercyyy"] = 10
		m := map[string]int{
			"Mercyyy":   10,
			"NoMercyyy": 20,
		}
		fmt.Println(m)		// m["Mercyyy"] = 10
	*/

	/*
		// zero value of map
		var m map[string]int // zero value
		fmt.Println(m == nil)
		m["Mercy"] = 10 //panic
	*/

	/*
		// map declaration using make
		m := make(map[string]int)
		m["Mercyyy"] = 10
		s := m["Mercyyy"]
		fmt.Println(m["Mercyyy"])
		fmt.Println(s)
	*/

	// check existence and read value
	m := map[string]int{
		"b": 10,
		"c": 20,
		"d": 30,
		"e": 40,
		"a": 50,
	}

	value, exist := m["d"]
	if exist {
		fmt.Println("values exists: ", value, exist)
	} else {
		fmt.Println("values  doesnt exists: ", value)
	}

	// delete key
	delete(m, "e")
	value, exist = m["e"]
	if exist {
		fmt.Println("values exists: ", value, exist)
	} else {
		fmt.Println("values  doesnt exists ")
	}

	// length
	fmt.Println("The length of map is", len(m))

	// range loop in map
	for key, value := range m {
		fmt.Println(key, value)
	}

	s := []string{}
	for i := range m {
		s = append(s, i)
		fmt.Println("range: ", s, i)
	}
	sort.Strings(s)

	for _, v := range s {
		fmt.Println("sorted:", v, m[v])
	}
	fmt.Println(s)
}
