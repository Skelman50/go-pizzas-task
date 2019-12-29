package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

type Pizza struct {
	Toppings []string `json:"toppings"`
}

type sortType struct {
	Toppings string
	Quantity int
}

func main() {
	pizzas := getJson()
	t1 := time.Now()
	pizzasMap := getPizzasMap(pizzas)
	result := getSortToppings(pizzasMap)
	fmt.Println(result)
	fmt.Println("elapsedTime:", time.Now().Sub(t1))
}

func getJson() []Pizza {
	jsonFile, err := os.Open("pizzas.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	pizzas := make([]Pizza, 0, 100000)
	json.Unmarshal(byteValue, &pizzas)
	if err != nil {
		fmt.Println(err)
	}
	return pizzas
}

func getPizzasMap(pizzas []Pizza) map[string]int {
	var pizzasMap = make(map[string]int)
	for i := 0; i < len(pizzas); i++ {
		sort.Strings(pizzas[i].Toppings)
		key := strings.Join(pizzas[i].Toppings, ", ")
		pizzasMap[key] += 1
	}
	return pizzasMap
}

func getSortToppings(pizzasMap map[string]int) []sortType {
	sortPizzasToStruct := make([]sortType, 0, 100000)
	for key, value := range pizzasMap {
		sortPizzasToStruct = append(sortPizzasToStruct, sortType{key, value})
	}
	sort.Slice(sortPizzasToStruct, func(i, j int) bool {
		return sortPizzasToStruct[i].Quantity > sortPizzasToStruct[j].Quantity
	})
	result := sortPizzasToStruct[0:20]
	return result
}
