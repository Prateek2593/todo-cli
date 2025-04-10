package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aquasecurity/table"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	Priority    string
}

type Todos []Todo

func (todos *Todos) Add(title string, priority string) {
	validPriorities := map[string]bool{
		"low":    true,
		"medium": true,
		"high":   true,
	}
	priority = strings.ToLower(priority)
	if !validPriorities[priority] {
		priority = "medium"
	}
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		Priority:    priority,
	}
	*todos = append(*todos, todo)
	fmt.Printf("Added todo: %s with priority: %s\n", title, priority)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("index out of range")
	}
	return nil
}
func (todos *Todos) Delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	todo := &t[index]
	todo.Completed = !todo.Completed
	if todo.Completed {
		now := time.Now()
		todo.CompletedAt = &now
	} else {
		todo.CompletedAt = nil
	}
	*todos = t
	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	todo := &t[index]
	todo.Title = title
	*todos = t
	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Index", "Title", "Completed", "Created At", "Completed At", "Priority")
	caser := cases.Title(language.English)
	for index, t := range *todos {
		completed := "No"
		completedAt := ""
		if t.Completed {
			completed = "Yes"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		fmt.Printf("Index: %d, Priority: %s\n", index, caser.String(t.Priority))
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt, caser.String(t.Priority))
	}
	table.Render()

}
