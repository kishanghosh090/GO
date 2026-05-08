package main

import "maps"

// maps -> hash, object, dict
func main() {
	// creating map key -> string  | value -> int
	// create map using make function

	m := make(map[string]int)

	m["key"] = 10
	m["area"] = 2323

	println(len(m))
	println(m["key"])

	// if key does not exists in the map it return zero value string - > '' bool -> false int -> 0
	println(m["dd"])

	delete(m, "area")

	clear(m)

	// new map
	newMap := map[string]int{"price": 23}

	println(newMap)

	// check is val exists ok -> true if have otherwise false
	_, ok := newMap["price"]

	if ok {
		println("all ok")
	}

	m1 := map[string]int{"price": 12, "phone": 3}
	m2 := map[string]int{"price": 12, "phone": 3}

	println(maps.Equal(m1, m2), m1 == nil)

	// empty map
	empMap := map[string]int{}
	if empMap == nil {

	}

}
