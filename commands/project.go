package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/dbalseiro/asana/api"
	"github.com/dbalseiro/asana/config"
	"github.com/dbalseiro/asana/utils"
)

func Projects(c *cli.Context) {
    projects := api.Projects()
    index := 0
    if len(projects) > 0 {
        fmt.Println("\n" + strconv.Itoa(len(projects)) + " projects found.")
        for i, p := range projects {
            fmt.Printf("[%d] %16d %s\n", i, p.Id, p.Name)
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


