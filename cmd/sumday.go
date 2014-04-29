package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/marcesher/sumday"
	"strings"
)

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

	//	RelMan	Qu	Team/General	Collab	CFG	Handbook	PTO/Holiday/Training	SES	20	WG
	tabs := []string{"rm", "qu", "team", "collab", "cfg", "handbook", "pto", "ses", "20", "wg"}
	cells := map[string]float64{}

	out := ""
	total := 0.0
	for _, cat := range tabs {
		cells[cat] = hours[cat]
		out += fmt.Sprintf("%v\t", hours[cat])
		total += hours[cat]
	}
	fmt.Println(strings.TrimSpace(out))
	fmt.Printf("%v Total Hours", total)
	clipboard.WriteAll(strings.TrimSpace(out))
}
