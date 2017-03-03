package commands

import (
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"

    "github.com/codegangsta/cli"

    "github.com/dbalseiro/asana/api"
    "github.com/dbalseiro/asana/utils"
)

func Description(c *cli.Context, withProject bool) {
    taskId := api.FindTaskId(c.Args().First(), false, withProject)
    task, _ := api.Task(taskId, true)

    tmpFile := os.TempDir() + "/asana_description.txt"
    f, err := os.Create(tmpFile)
    utils.Check(err)
    defer f.Close()

    err = templateDescription(f, task)
    utils.Check(err)

    cmd := exec.Command(os.Getenv("EDITOR"), tmpFile)
    cmd.Stdin, cmd.Stdout = os.Stdin, os.Stdout
    err = cmd.Run()

    txt, err := ioutil.ReadFile(tmpFile)

    utils.Check(err)

    postDescription := trim(string(txt))

    t := api.Update(taskId, "notes", postDescription)

    fmt.Println("Description on Task: \"" + t.Name + "\"\n")
    fmt.Println(t.Notes)
}

func templateDescription(f *os.File, task api.Task_t) error {
    var err error
    _, err = f.WriteString(task.Notes)
    return err
}
