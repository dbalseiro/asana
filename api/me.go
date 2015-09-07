package api

import (
	"encoding/json"
	"strconv"
	"github.com/dbalseiro/asana/config"
	"github.com/dbalseiro/asana/utils"
)

type User_t struct {
	Id         int
	Name       string
	Email      string
	Workspaces []Base
	Photo      map[string]string
}

type Me_t User_t

func Me() Me_t {
	var me map[string]Me_t
	err := json.Unmarshal(Get("/api/1.0/users/me", nil), &me)
	utils.Check(err)
	return me["data"]
}

func FindUserId(name string) string {
    var users map[string][]User_t
    uri := "/api/1.0/workspaces/" + strconv.Itoa(config.Load().Workspace) + "/typeahead?type=user&query=" + name
    err := json.Unmarshal(Get(uri, nil), &users)
    utils.Check(err)
    return strconv.Itoa(users["data"][0].Id)
}
