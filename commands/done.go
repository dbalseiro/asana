package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/dbalseiro/asana/api"
)

func Done(c *cli.Context, withProject bool) {
	task := api.Update(api.FindTaskId(c.Args().First(), false, withProject), "completed", "true")
	fmt.Println("DONE! : " + task.Name)
}
