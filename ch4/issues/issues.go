package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"../github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d тем: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %-55.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt.Format(time.RFC3339))
	}
}
