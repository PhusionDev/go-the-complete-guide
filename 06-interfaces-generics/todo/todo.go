package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	text string
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("invalid input")
	}

	return &Todo{
		text,
	}, nil
}

func (t *Todo) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Text string `json:"text"`
	}{
		Text: t.text,
	})
}

func (t *Todo) Print() {
	fmt.Printf("Todo: %s\n", t.text)
}

func (t *Todo) Save() error {
	fileName := "todo.json"

	data, err := t.MarshalJSON()
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
