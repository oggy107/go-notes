package main

import (
	"fmt"
)

func mapsExample() {
	// ages := map[string]int{
	// 	"john":  19,
	// 	"james": 23,
	// }

	ages := make(map[string]int)
	ages["john"] = 87
	ages["smith"] = 69

	fmt.Println(ages)
	fmt.Println(len(ages))

	// inserting
	ages["oggy"] = 22

	// getting
	// age := ages["smith"]

	// deleting
	delete(ages, "smith")

	// check if key exists
	s, ok := ages["smith"]
	fmt.Println(s, ok)
}
