package main

import (
	"os"
	"github.com/codegangsta/cli"

	"github.com/dbalseiro/asana/commands"
	"github.com/dbalseiro/asana/config"
)

func main() {
	app := cli.NewApp()
	app.Name = "asana"
	app.Version = "0.1.2"
	app.Usage = "asana cui client"

	app.Commands = defs()
	app.Run(os.Args)
}

func isWithProject() bool {
    p := config.Load().Project
    return p != 0
}

func defs() []cli.Command {
	return []cli.Command{
		{
			Name:      "config",
			ShortName: "c",
			Usage:     "Asana configuration. Your settings will be saved in ~/.asana.yml",
			Action: func(c *cli.Context) {
				commands.Config(c)
			},
		},
		{
			Name:      "workspaces",
			ShortName: "w",
			Usage:     "get workspaces",
			Action: func(c *cli.Context) {
				commands.Workspaces(c)
			},
		},
        {
			Name:      "project-tasks",
			ShortName: "pt",
			Usage:     "get project tasks",
			Action: func(c *cli.Context) {
				commands.Tasks(c, true)
			},
		},
{
			Name:      "projects",
			ShortName: "p",
			Usage:     "get workspaces",
			Action: func(c *cli.Context) {
				commands.Projects(c)
			},
		},
		{
			Name:      "tasks",
			ShortName: "ts",
			Usage:     "get tasks",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "no-cache, n", Usage: "without cache"},
				cli.BoolFlag{Name: "refresh, r", Usage: "update cache"},
			},
			Action: func(c *cli.Context) {
				commands.Tasks(c, false)
			},
		},
		{
			Name:      "task",
			ShortName: "t",
			Usage:     "get a task",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "verbose, v", Usage: "verbose output"},
			},
			Action: func(c *cli.Context) {
				commands.Task(c, isWithProject())
			},
		},
		{
			Name:      "comment",
			ShortName: "cm",
			Usage:     "Post comment",
			Action: func(c *cli.Context) {
				commands.Comment(c, isWithProject())
			},
		},
		{
			Name:      "done",
			Usage:     "Complete task",
			Action: func(c *cli.Context) {
				commands.Done(c, isWithProject())
			},
		},
		{
			Name:  "due",
			Usage: "set due date",
			Action: func(c *cli.Context) {
				commands.DueOn(c, isWithProject())
			},
		},
        {
			Name:  "assign",
			Usage: "assign task",
			Action: func(c *cli.Context) {
				commands.Assign(c, isWithProject())
			},
		},
	}
}
