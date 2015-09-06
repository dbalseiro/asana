package api

import (
    "strconv"
    "encoding/json"
	"github.com/dbalseiro/asana/config"
	"github.com/dbalseiro/asana/utils"
)

type Project_t struct {
    Id int
    Name string
}

func Projects() []Project_t {
    var projects map[string][]Project_t
    uri := "/api/1.0/workspaces/" + strconv.Itoa(config.Load().Workspace) + "/projects"
    err := json.Unmarshal(Get(uri, nil), &projects)
    utils.Check(err)
    return projects["data"]
}


