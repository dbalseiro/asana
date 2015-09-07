package commands

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"github.com/codegangsta/cli"

	"github.com/dbalseiro/asana/api"
	"github.com/dbalseiro/asana/utils"
)

const (
	CacheDuration = "5m"
)

func Tasks(c *cli.Context, withProject bool) {
	if c.Bool("no-cache") {
		fromAPI(false, withProject)
	} else {
		if utils.Older(CacheDuration, utils.CacheFile()) || c.Bool("refresh") {
			fromAPI(true, withProject)
		} else {
			txt, err := ioutil.ReadFile(utils.CacheFile())
			if err == nil {
				lines := regexp.MustCompile("\n").Split(string(txt), -1)
				for _, line := range lines {
                    if line != "" {
                        format(line)
                    }
				}
			} else {
				fromAPI(true, withProject)
			}
		}
	}
    if !withProject {
        ClearConfig()
    }
}

func fromAPI(saveCache bool, withProject bool) {
	tasks := api.Tasks(url.Values{}, false, withProject)
	if saveCache {
		cache(tasks)
	}
	for i, t := range tasks {
        fmt.Printf("%2d [ %10s ] @%s: %s ", i, t.Due_on, t.Assignee.Name, t.Name)
        for _, p := range t.Projects {
            fmt.Printf("#%s ", p.Name)
        }

        for _, ta := range t.Tags {
            fmt.Printf("#%s ", ta.Name)
        }
        fmt.Printf("\n")
	}
}

func cache(tasks []api.Task_t) {
	f, _ := os.Create(utils.CacheFile())
	defer f.Close()
	for i, t := range tasks {
		f.WriteString(strconv.Itoa(i) + ":")
		f.WriteString(strconv.Itoa(t.Id) + ":")
		f.WriteString(t.Due_on + ":")
        f.WriteString(t.Assignee.Name + ":")
		f.WriteString(t.Name + " ")
        for _, p := range t.Projects {
            f.WriteString("#" + p.Name + " ")
        }

        for _, ta := range t.Tags {
            f.WriteString("#" + ta.Name + " ")
        }
        f.WriteString("\n")
}
}

func format(line string) {
	dateRegexp := "[0-9]{4}-[0-9]{2}-[0-9]{2}"

	index := regexp.MustCompile("^[0-9]*").FindString(line)
	line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove index
	line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove task_id
	
    date := regexp.MustCompile("^" + dateRegexp).FindString(line)
	line = regexp.MustCompile("^("+dateRegexp+")?:").ReplaceAllString(line, "") // remove date
    
    assignee := regexp.MustCompile("^[a-zA-Z]*:").FindString(line)
	line = regexp.MustCompile("^.*:").ReplaceAllString(line, "") // remove assignee

	fmt.Printf("%2s [ %10s ] @%s %s\n", index, date, assignee, line)
}
