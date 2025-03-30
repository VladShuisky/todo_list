package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type TodoList struct {
	Todos []string `json:"todo_list"`
}

func readTodoList(filename string) (TodoList, error) {
	var todoList TodoList
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return TodoList{Todos: []string{}}, nil
		}
		return todoList, err
	}
	err = json.Unmarshal(data, &todoList)
	return todoList, err
}

func saveTodoList(filename string, todoList TodoList) error {
	data, err := json.MarshalIndent(todoList, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func addTodo(filename, newTodo string) error {
	todoList, err := readTodoList(filename)
	if err != nil {
		return err
	}
	todoList.Todos = append(todoList.Todos, newTodo)
	return saveTodoList(filename, todoList)
}

func showTodoList(filename string) {
	todoList, err := readTodoList(filename)
    if err != nil {
        fmt.Println("Ошибка чтения:", err)
        return
    }

    fmt.Println("Текущий список дел:")
    for _, todo := range todoList.Todos {
        fmt.Println(todo)
    }
}

func main() {
	filename := "todo_list.json"
	if len(os.Args) <= 1 {
		showTodoList(filename)
		return
	}
	if os.Args[1] == "add" {
		new_todo_string := os.Args[2]
		todo_list, err := readTodoList(filename)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			return
		}
		new_todo_index := strconv.Itoa(len(todo_list.Todos)) + ". "
		addTodo(filename, new_todo_index + new_todo_string)
	}
	showTodoList(filename)
}