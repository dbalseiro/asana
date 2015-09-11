package commands

import (
	"fmt"
    "strings"
    "regexp"

	"github.com/codegangsta/cli"
    "github.com/mgutz/ansi"

	"github.com/dbalseiro/asana/api"
)

func CreateTask(c *cli.Context, withProject bool) {
    task := api.CreateTask(withProject, c.String("name"))

    fmt.Println(task)
}

func Task(c *cli.Context, withProject bool) {
	t, stories := api.Task(api.FindTaskId(c.Args().First(), true, withProject), c.Bool("verbose"))
    red := ansi.ColorCode("red")
    blue := ansi.ColorCode("blue")
    green := ansi.ColorCode("green")
    reset := ansi.ColorCode("reset")
    bold := ansi.ColorCode("white")
    yellow := ansi.ColorCode("yellow")

	fmt.Printf("[ %s%s%s ] %s%s%s\n", red, t.Due_on, reset, bold, t.Name, reset)

	showTags(t.Tags)

	fmt.Printf("\n%s\n", t.Notes)

	if stories != nil {
		fmt.Println("\n----------------------------------------\n")
		for _, s := range stories {
            color := ""

            if strings.HasPrefix(s.String(), "*") {
                color = blue
            }
            match, _ := regexp.MatchString("\\* .* attached", s.String())
            if match {
                color = yellow
            }

            if strings.HasPrefix(s.String(), ">") {
                color = green
            }

			fmt.Printf("%s%s%s\n", color, s, reset)
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
