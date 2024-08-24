package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 20
	ages["charlie"] = 30
	fmt.Println(ages)

	delete(ages, "charlie")
	fmt.Println(ages)

	ages = map[string]int{
		"alice":   18,
		"charlie": 32,
	}
	fmt.Println(ages)

	ages["bob"] = 40
	fmt.Println(ages)

	ages["bob"] += 1
	fmt.Println(ages)

	ages["charlie"]++
	fmt.Println(ages)

	fmt.Printf("Len = %d\n", len(ages))

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	age, ok := ages["hung"]
	if !ok {
		fmt.Println("hung is not a key in this map")
	}
	fmt.Printf("Age = %d\n", age)
}
