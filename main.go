package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	title       string
	description string
}

type TodoContainer struct {
	items []Todo
	index int
}

type OptionError struct {
}

func (e *OptionError) Error() string {
	return "Invalid option"
}

func (t *Todo) String() string {
	return fmt.Sprintf("title: %v, description: %v", t.title, t.description)
}

func createTodo(title string, description string) Todo {
	return Todo{title: title, description: description}
}

func printOpts() {
	fmt.Println("\nOptions")
	fmt.Println("  1. Create new todo")
	fmt.Println("  2. List todos")
	fmt.Println("  3. Save todos")
	fmt.Println("  4. Load todos")
	fmt.Println("  5. Quit")
}

func getSelection() (int, error) {
	var option int
	_, err := fmt.Scan(&option)

	if err != nil || option > 5 {
		return -1, &OptionError{}
	}

	return option, err
}

func handleSelection(selection int, todos *TodoContainer) {
	switch selection {
	case 1:
		addTodo(todos)
	case 2:
		listTodos(todos)
	case 3:
		saveTodos(todos)
	case 4:
		readTodos(todos)
	case 5:
		os.Exit(0)
	default:
		break
	}
}

func addTodo(todos *TodoContainer) {
	fmt.Println("\nAdd title for new todo")
	title := readLine()

	fmt.Println("Add description for new todo")
	description := readLine()

	todo := createTodo(title, description)
	fmt.Println("New todo created: ", todo.String())

	todos.items = append(todos.items, todo)
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error while reading input")
	}

	input = strings.TrimSuffix(input, "\n")
	return input
}

func listTodos(todos *TodoContainer) {
	fmt.Println("\nAll todos:")
	for _, item := range todos.items {
		fmt.Println(item.String())
	}
	fmt.Println()
}

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

func loop() {
	todos := &TodoContainer{items: make([]Todo, 0)}

	for {
		printOpts()
		selection, err := getSelection()

		if err != nil {
			fmt.Println(err)
			break
		}

		handleSelection(selection, todos)
	}
}

func main() {
	loop()
}
