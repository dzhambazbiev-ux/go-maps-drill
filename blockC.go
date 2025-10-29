package main

import "fmt"


//func main() {
	students := map[int]string {
		1:"Dzhambulat",
		2:"Deni",
		3:"Vaha", 
		4:"Muslim",
	}
	fmt.Println(students, len(students))

	delete(students, 2)
	fmt.Println(students, len(students))
	fmt.Println(students[2])

	for k, val := range students {
		if len(val) != 0 {
			delete(students, k)
		}
	}
	fmt.Println("Длина мапа:", len(students))

	delete(students, 2)
	fmt.Println(students[2])
}