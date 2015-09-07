package commands

import (
	"fmt"
	"os"
	"strconv"

    "github.com/mgutz/ansi"
	"github.com/codegangsta/cli"
	"github.com/dbalseiro/asana/api"
	"github.com/dbalseiro/asana/config"
	"github.com/dbalseiro/asana/utils"
)

func Projects(c *cli.Context) {
    projects := api.Projects()
    index := 0
    cyan := ansi.ColorCode("cyan")
    reset := ansi.ColorCode("reset")
    if len(projects) > 0 {
        fmt.Println("\n" + strconv.Itoa(len(projects)) + " projects found.")
        for i, p := range projects {
            fmt.Printf("%s[%2d]%s %s\n", cyan, i, reset, p.Name)
        }
        index = utils.EndlessSelect(len(projects)-1, index)
    }
    
    workspace := config.Load().Workspace
    apiKey := config.Load().Api_key
    proj := strconv.Itoa(projects[index].Id) 

    f, _ := os.Create(utils.Home() + "/.asana.yml")
    f.WriteString("api_key: " + apiKey + "\n")
    f.WriteString("workspace: " + strconv.Itoa(workspace) + "\n")
    f.WriteString("project: " + proj + "\n")
}


