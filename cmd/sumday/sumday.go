package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/marcesher/sumday"
)

type Category struct {
	Name    string
	Aliases []string
}

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	hours := sumday.SumDay(text)
	fmt.Println(hours)

	if len(hours) == 0 {
		fmt.Printf("YOOOO!!!!! No hours present in input string \n\n %v \n\n", text)
	}

	categories, err := ioutil.ReadFile("config.json")
	var cats []Category
	err = json.Unmarshal(categories, &cats)
	if err != nil {
		panic(err)
	}

	out := ""
	total := 0.0

	for _, cat := range cats {
		cat_total := 0.0
		cat_total = hours[cat.Name]
		for _, alias := range cat.Aliases {
			cat_total += hours[alias]
			delete(hours, alias)
		}
		delete(hours, cat.Name)
		out += fmt.Sprintf("%v\t", cat_total)
		total += cat_total
	}

	fmt.Println(strings.TrimSpace(out))
	fmt.Printf("%v Total Hours\n", total)
	fmt.Printf("Uncategorized hours is %v\n", hours)
	clipboard.WriteAll(strings.TrimSpace(out))
}
