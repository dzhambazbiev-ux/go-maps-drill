package main

import "fmt"

func main() {
	var m map[string]int 
	//m["x"] = 1  // нельзя объявить мап и не инициализировать ее черeз {}, либо через make()
	m = make(map[string]int)
	m["x"] = 1
	fmt.Println(m)

	// 2
	// bad := map[[]int]string{} // срез нельзя писать в ключ мапа, потому что срез это не сравнимы тип данных
	// fmt.Println(bad)

	fmt.Println("")

	// 3
	roomX := map[string]bool{}
	fmt.Println(roomX["room"]) // вернет false,  потому ключа нет, нулевое значение bool = false

	v, ok := roomX["room"]
	fmt.Println(v, ok) // вернет false false, первый false нулевое значение, второй false потому что нет такого ключа

}
