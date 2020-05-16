// Упражнение 4.10 - группировка результата по группам:
// "За текущий месяц" и "За текущий год"
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"../github"
)

func main() {
	// Создаю два среза с типом стрктуры хранения Issues.
	// Этоа структура описана в пакете github.
	var lastMonth []*github.Issues
	var lastYear []*github.Issues

	// Беру текущий год и месяц
	yearNow, monthNow, _ := time.Now().Date()

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		yearCreated, monthCreated, _ := item.CreatedAt.Date()

		// Группировка по признаку "За текущий месяц"
		if yearNow == yearCreated && monthCreated-monthNow == 0 {
			// Добавляю текущий Item в срез lastMonth
			lastMonth = append(lastMonth, item)
		}
		// Группировка по признаку "За текущий год"
		if yearNow == yearCreated {
			// Добавляю текущий Item в срез lastYear
			lastYear = append(lastYear, item)
		}
	}

	fmt.Printf("%d тем за всё время.\n", result.TotalCount)

	fmt.Printf("[Тут количество] тем в текущем месяце:\n")
	for _, item := range lastMonth {
		fmt.Printf("#%-5d %9.9s %-55.55s %10.10s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Printf("\n[Тут количество] тем в текущем году:\n")
	for _, item := range lastYear {
		fmt.Printf("#%-5d %9.9s %-55.55s %10.10s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

}
