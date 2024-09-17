package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title. id: new_title")
	flag.IntVar(&cf.Del, "delete", -1, "Specify a todo index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute (todos *Todos) {
	switch {
	case cf.List:
		todos.print()

	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error! Invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error! Invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")
		fmt.Println("Enter commands mentioned below:\ngo run ./ -list\t\t\t\t\t\t-- to view all items\ngo run ./ -add '<add_title>'\t\t\t\t-- to add item\ngo run ./ -edit '<item_number: new_title>'\t\t-- to edit item\ngo run ./ -toggle <item_number>\t\t\t\t--to toggle between item\ngo run ./ -delete <item_number>\t\t\t\t--to delete item")
	}
}
