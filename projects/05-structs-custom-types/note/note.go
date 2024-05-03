package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("invalid input")
	}

	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (n *Note) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}{
		Title:     n.title,
		Content:   n.content,
		CreatedAt: n.createdAt,
	})
}

func (n *Note) PrintNote() {
	fmt.Printf("Title: %s\nContent: %s\n", n.title, n.content)
}

func (n *Note) PrintTitle() {
	fmt.Println("Note Title:", n.title)
}

func (n *Note) PrintContent() {
	fmt.Println("Note Content:", n.content)
}

func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	data, err := n.MarshalJSON()

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
