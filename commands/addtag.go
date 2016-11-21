package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/dbalseiro/asana/api"
)

func AddTag(c *cli.Context, withProject bool) {
    tagname := c.Args()[1]
    tag := api.GetTag(tagname)
    taskId := api.FindTaskId(c.Args().First(), false, false)

    api.AddTagToTask(tag.Id, taskId)
	fmt.Println("assigned! : " + tag.Name + " to " + taskId)
}
