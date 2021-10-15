package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func saveTodos(todos *TodoContainer) {
	f, err := os.Create("todos.csv")

	if err != nil {
		fmt.Println("Error while creating csv file")
	}

	defer f.Close()

	fmt.Fprintln(f, "title,description")

	for _, line := range todos.items {
		_, err := fmt.Fprintf(f, "%v,%v\n", line.title, line.description)
		if err != nil {
			fmt.Println("Error with line ", err)
		}
	}

	fmt.Println("\nTodos saved")
}

func readTodos(todos *TodoContainer) {
	f, err := os.Open("todos.csv")

	if err != nil {
		fmt.Println("Couldn't open todo file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Skip first row
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		title, description := parts[0], parts[1]
		todos.items = append(todos.items, createTodo(title, description))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
