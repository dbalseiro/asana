package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/codegangsta/cli"

	"github.com/dbalseiro/asana/api"
	"github.com/dbalseiro/asana/utils"
)

func Comment(c *cli.Context, withProject bool) {
	taskId := api.FindTaskId(c.Args().First(), false, withProject)
	task, stories := api.Task(taskId, true)

	tmpFile := os.TempDir() + "/asana_comment.txt"
	f, err := os.Create(tmpFile)
	utils.Check(err)
	defer f.Close()

	err = template(f, task, stories)
	utils.Check(err)

	cmd := exec.Command(os.Getenv("EDITOR"), tmpFile)
	cmd.Stdin, cmd.Stdout = os.Stdin, os.Stdout
	err = cmd.Run()

	txt, err := ioutil.ReadFile(tmpFile)

	utils.Check(err)

    isForClose := getIsForClose(string(txt))
    asignee := getAsignee(string(txt))
	postComment := trim(string(txt))

	if postComment != "" {
		commented := api.CommentTo(taskId, postComment)
		fmt.Println("Commented on Task: \"" + task.Name + "\"\n")
		fmt.Println(commented)

        if isForClose {
            result := api.Update(taskId, "completed", "true")
            fmt.Println("Task closed \"" + task.Name + "\"\n")
            fmt.Println(result)
        }

        if asignee != "" {
            result := api.Update(taskId, "assignee", api.FindUserId(strings.Replace(asignee, "@", "",-1)))
            fmt.Println("New asignee \"" + asignee + "\"\n")
            fmt.Println(result)
        }
    } else {
		fmt.Println("Aborting comment due to empty content.")
	}
}

func getIsForClose(txt string) bool {
    result := strings.Split(txt, "\n")
    for i := range result {
        if result[i] == "%CLOSE" {
            return true
        }
    }
    return false
}

func getAsignee(txt string) string {
    result := strings.Split(txt, "\n")
    for i := range result {
        if strings.HasPrefix(result[i], "@") {
            return result[i]
        }
    }
    return ""
}

func template(f *os.File, task api.Task_t, stories []api.Story_t) error {
	var err error
	_, err = f.WriteString("\n\n\n")
	_, err = f.WriteString("# =================================== \n")
	_, err = f.WriteString("# " + task.Name + "\n#\n")
	_, err = f.WriteString(commentOut(task.Notes) + "\n#\n")
	_, err = f.WriteString("\n# ----------------------------------- \n")
	for _, s := range stories {
		_, err = f.WriteString(commentOut(fmt.Sprintf("%s", s)) + "\n")
	}
	return err
}

func commentOut(txt string) string {
	return strings.Replace("# "+txt, "\n", "\n# ", -1)
}

func trim(txt string) string {
	var result string
	result = regexp.MustCompile("\n@.*\n").ReplaceAllString(txt, "")    // Remove asignees
	result = regexp.MustCompile("\n%CLOSE\n").ReplaceAllString(result, "")    // Remove close mark
	result = regexp.MustCompile("#.*\n").ReplaceAllString(result, "")    // Remove comments
	result = regexp.MustCompile("\n*$").ReplaceAllString(result, "")  // Remove blank lines
	result = regexp.MustCompile("\n").ReplaceAllString(result, "\\n") // Escape
	result = regexp.MustCompile("\"").ReplaceAllString(result, "\\\"") // comillas
	return result
}
