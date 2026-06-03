package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input(label string) string {
	fmt.Println(label)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Question(label string, options []string) []int {
	question := label
	for index, option := range options {
		question += fmt.Sprintf("\n%d) %s", index+1, option)
	}
	question += "\nEnter choices (comma separated, e.g. 1,2): "
	fmt.Print(question)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	choices := []int{}
	for _, part := range strings.Split(line, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		answer, err := strconv.Atoi(part)
		if err != nil || answer < 1 || answer > len(options) {
			fmt.Println("Invalid choice, please try again.")
			return Question(label, options)
		}
		choices = append(choices, answer-1)
	}

	if len(choices) == 0 {
		fmt.Println("No choice made, please try again.")
		return Question(label, options)
	}

	return choices
}

func QuestionSingle(label string, options []string) int {
	question := label
	for index, option := range options {
		question += fmt.Sprintf("\n%d) %s", index+1, option)
	}
	question += fmt.Sprintf("\nEnter choice (1-%d): ", len(options))
	fmt.Print(question)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	answer, err := strconv.Atoi(input)
	if err != nil || answer < 1 || answer > len(options) {
		fmt.Println("Invalid choice, please try again.")
		return QuestionSingle(label, options)
	}
	return answer - 1
}

func WaitForEnter() {
	fmt.Println("\nPress Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func ShowErrorAndExit(errMsg string) {
	fmt.Println("\n==========================================")
	fmt.Printf("ERROR: %s\n", errMsg)
	fmt.Println("==========================================")
	WaitForEnter()
	os.Exit(1)
}
