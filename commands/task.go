package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/dbalseiro/asana/api"
)

func CreateTask(c *cli.Context, withProject bool) {
    task := api.CreateTask(withProject, c.String("name"))

    fmt.Println(task)
}

func Task(c *cli.Context, withProject bool) {
	t, stories := api.Task(api.FindTaskId(c.Args().First(), true, withProject), c.Bool("verbose"))

	fmt.Printf("[ %s ] %s\n", t.Due_on, t.Name)

	showTags(t.Tags)

	fmt.Printf("\n%s\n", t.Notes)

	if stories != nil {
		fmt.Println("\n----------------------------------------\n")
		for _, s := range stories {
			fmt.Printf("%s\n", s)
		}
	}
}

func showTags(tags []api.Base) {
	if len(tags) > 0 {
		fmt.Print("  Tags: ")
		for i, tag := range tags {
			print(tag.Name)
			if len(tags) != 1 && i != (len(tags)-1) {
				print(", ")
			}
		}
		println("")
	}
}
