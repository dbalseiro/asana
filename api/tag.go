package api

import (
	"encoding/json"
	"strconv"

	"github.com/dbalseiro/asana/config"
	"github.com/dbalseiro/asana/utils"
)

type Tags_t struct {
    Data []Tag_t
}

type Tag_t struct {
	Id              int
	Name            string
}

func Tags() Tags_t {
    var tags Tags_t
    uri := "/api/1.0/workspaces/" + strconv.Itoa(config.Load().Workspace) + "/tags"

    err := json.Unmarshal(Get(uri, nil), &tags)
    utils.Check(err)

    return tags
}

func GetTag(tagname string) Tag_t {
    for _, t := range Tags().Data {
        if t.Name == tagname {
            return t
        }
    }
    return CreateTag(tagname)
}

func CreateTag(tagname string) Tag_t {
    workspace := strconv.Itoa(config.Load().Workspace)
    data := `{
        "data": {
            "workspace": ` + workspace + `,
            "name": "` + tagname + `"
        }
    }`

    respBody := Post("/tags", data)
    var output map[string]Tag_t
    err := json.Unmarshal(respBody, &output)
    utils.Check(err)

    newTag := output["data"]
    return newTag
}

func AddTagToTask(tagId int, taskId string) {
    uri := "/tasks/" + taskId + "/addTag";
    data := `{
        "data": {
            "tag": ` + strconv.Itoa(tagId) + `
        }
    }`
    Post(uri, data)
}
