package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add      string
	Delete   int
	Edit     string
	Toggle   int
	List     bool
	Priority string
	Sort     bool
	Notes    string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title. id:new_title")
	flag.IntVar(&cf.Delete, "del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.StringVar(&cf.Priority, "p", "medium", "Specify a priority for a todo")
	flag.BoolVar(&cf.Sort, "sort", false, "Sort todos by priority")
	flag.StringVar(&cf.Notes, "notes", "", "Specify notes for a todo")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print(cf.Sort)
	case cf.Add != "":
		todos.Add(cf.Add, cf.Priority, cf.Notes)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit command. Use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index:", parts[0])
			os.Exit(1)
		}
		todos.Edit(index, parts[1])
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	default:
		fmt.Println("Invalid command")
	}

}
