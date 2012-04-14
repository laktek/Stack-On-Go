package main

import "fmt"
import "github.com/laktek/stack-on-go/stackongo"

func main() {
	fmt.Printf("Here are some facts about StackOverflow:\n")

	session := stack-on-go.NewSession("stackoverflow")
	info, err := session.Info()

	if err != nil {
		fmt.Printf(err.String())
	}

	fmt.Printf("Total Questions: %v\n", info.Total_questions)
	fmt.Printf("Total Answers: %v\n", info.Total_answers)
	fmt.Printf("Total Users: %v\n", info.Total_users)

	fmt.Printf("Questions per minute: %v\n", info.Questions_per_minute)
	fmt.Printf("Answers per minute: %v\n", info.Answers_per_minute)
	fmt.Printf("New active users: %v\n", info.New_active_users)
}
