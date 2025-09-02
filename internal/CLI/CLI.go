package cli

import (
	"fmt"
	storage "ivan/CLI-Task-Tracker/internal/Storage"
	task "ivan/CLI-Task-Tracker/internal/Task"
	"os"
	"strconv"
)

type CLI struct {
	storage storage.Storage
}

func NewCLI(storage storage.Storage) *CLI {
	return &CLI{
		storage: storage,
	}
}

func (c *CLI) Run() error {
	if len(os.Args) < 2 {
		c.printFunctions()
		return nil
	}

	cmd := os.Args[1]

	switch cmd {
	case "list":
		tasks, err := c.storage.List()
		if err != nil {
			return err
		}
		for _, t := range tasks {
			fmt.Println("==========")
			fmt.Println("id: ", t.ID)
			fmt.Println("description: ", t.Description)
			fmt.Println("status: ", t.Status)
		}

	case "add":
		if len(os.Args) < 3 {
			return fmt.Errorf("too few args")
		}

		var description string

		for i := 2; i < len(os.Args); i++ {
			description = fmt.Sprint(description, " ", os.Args[i])
		}

		if err := c.storage.Add(*task.NewTask(description)); err != nil {
			return err
		}

	case "done":
		if len(os.Args) < 3 {
			return fmt.Errorf("too few args")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return err
		}

		err = c.storage.SetTaskStatusComplited(id)
		if err != nil {
			return err
		}

	case "get":
		if len(os.Args) < 3 {
			return fmt.Errorf("too few args")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return err
		}

		task, err := c.storage.GetByID(id)
		if err != nil {
			return err
		}

		fmt.Println("id: ", task.ID)
		fmt.Println("description: ", task.Description)
		fmt.Println("status: ", task.Status)

	case "remove":
		if len(os.Args) < 3 {
			return fmt.Errorf("too few args")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return err
		}

		if err = c.storage.Delete(id); err != nil {
			return err
		}

	case "update":
		if len(os.Args) < 3 {
			return fmt.Errorf("too few args")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return err
		}

		var newDescription string

		for i := 2; i < len(os.Args); i++ {
			newDescription = fmt.Sprint(newDescription, " ", os.Args[i])
		}

		if err = c.storage.Update(id, newDescription); err != nil {
			return err
		}

	case "completed":
		tasks, err := c.storage.GetComplitedTasks()
		if err != nil {
			return err
		}

		for _, t := range tasks {
			fmt.Println("==========")
			fmt.Println("id: ", t.ID)
			fmt.Println("description: ", t.Description)
			fmt.Println("status: ", t.Status)
		}

	case "pending":
		tasks, err := c.storage.GetPendingTasks()
		if err != nil {
			return err
		}

		for _, t := range tasks {
			fmt.Println("==========")
			fmt.Println("id: ", t.ID)
			fmt.Println("description: ", t.Description)
			fmt.Println("status: ", t.Status)
		}

	default:
		c.printFunctions()
	}

	return nil
}

func (c *CLI) printFunctions() {
	fmt.Println("list")
	fmt.Println("add <description>")
	fmt.Println("done <id>")
	fmt.Println("get <id>")
	fmt.Println("remove <id>")
	fmt.Println("update <id> <new description>")
	fmt.Println("completed")
	fmt.Println("pending")
}
