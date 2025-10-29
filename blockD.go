package main

import (
	"fmt"
	"sort"
)

//func main() {
	loc := map[string]string{
		"Грозный":     "Мечеть сердце Чечни",
		"Урус-Мартан": "Авторынок",
		"Казеной":     "Озеро",
		"Закан-юрт":   "Природа",
		"Астрахань":   "Площадь",
	}

	for k, val := range loc {
		fmt.Printf("%s - %s\n", k, val)
	}

	// порядок в map не важен, важно ключ и его значение, map это ассоциативный список

	// 3
	for k, _ := range loc {
		fmt.Println(k)
	}

	for _, val := range loc {
		fmt.Println(val)
	}

	// 4
	slice := []string{}
	for k, _ := range loc {
		slice = append(slice, k)
	}
	sort.Strings(slice)

	
	for i, val := range slice {
		fmt.Printf("%d. %s\n", i+1, val)
	}

	// 5
	for k, val := range loc {
		fmt.Printf("%s: %s\n", k, val)
	}
}
