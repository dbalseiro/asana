package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/dbalseiro/asana/api"
)

func Assign(c *cli.Context, withProject bool) {
	task := api.Update(api.FindTaskId(c.Args().First(), false, withProject), "assignee", c.Args()[1])
	fmt.Println("assigned! : " + task.Name + " to " + task.Assignee.Name)
}

func AssignMe(c *cli.Context, withProject bool) {
	task := api.Update(api.FindTaskId(c.Args().First(), false, withProject), "assignee", "me")
	fmt.Println("assigned! : " + task.Name + " to " + task.Assignee.Name)
}
