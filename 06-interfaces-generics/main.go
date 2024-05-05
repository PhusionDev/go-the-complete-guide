package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

// type printer interface {
// 	Print()
// }

type outputtable interface {
	saver
	Print()
}

// type outputtable interface {
// 	Save() error
// 	Print()
// }

func main() {
	title, content := getNoteData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)
	if err != nil {
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}
}

func printSomething(value any) {
	typedVal, ok := value.(int)
	if ok {
		typedVal += 1
	}
	// switch value.(type) {
	// case int:
	//   fmt.Println("Integer:", value)
	// case float64:
	//   fmt.Println("Float:", value)
	// default:
	//   fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Print()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data has failed...")
		return err
	}

	fmt.Println("Saving the data was successful...")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
