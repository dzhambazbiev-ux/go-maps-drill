package main

import "fmt"

//func main() {
	week := map[string]int{
		"Понедельник": 2,
		"Вторник":     4,
		"Среда":       7,
	}

	fmt.Println(week, len(week))

	id := map[string]string {
		"akhmad":"Кудузов Ахмад",
		"rasul":"Назиров Расул",
	}
	fmt.Println(id["akhmad"])

	floor := map[int]string {
		1:"кухня",
		2:"коворкинг",
	}
	fmt.Println(floor[3])

	exchangeRate := map[string]float64 {
		"dollar":79.12,
		"euro": 92.43,
	}
	exchangeRate["dollar"] = 77.31
	fmt.Println(exchangeRate)

	loc := map[string]string {
		"1":"Мечеть сердце Чечни",
		"2":"Высотки Грозный Сити",
		"3":"Триумфальная Арка 'Грозный'",
	}
	
	b := loc
	b["3"] = "Грозный Молл"
	fmt.Println(loc, b)

	// 7

	newMap := map[string]string {
		"1":"America",
		"2":"Europe",
	}
	newMap["2"] = "Asia"

	fmt.Println(len(newMap))

	// 8 
	emptyMap := map[string]string {
		"":"hello",
		"hi":"",
	}
	fmt.Println(emptyMap)

	// block b 1
	fmt.Println(GetName(id,"akhmad"))

	// 2
	fmt.Println(HasFloor(floor, 1))

	// 3
	fmt.Println(ReadPrice(exchangeRate, "dollar"))
}

// block B
func GetName(m map[string]string, id string) (string, bool) {
	v, ok := m[id]
	return  v, ok
}

func HasFloor(m map[int]string, floor int) bool {
	_, ok := m[floor]

	return  ok
}

func ReadPrice(m map[string]float64, course string) (float64, bool) {
	val, ok := m[course]

	if ok {
		fmt.Println("Цена курса:", val)
	} else{
		fmt.Println("Курса нет")
	}
	return  val, ok
}