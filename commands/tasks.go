package commands

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strconv"

    "github.com/mgutz/ansi"
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
    red := ansi.ColorCode("red")
    blue := ansi.ColorCode("blue")
    green := ansi.ColorCode("green")
    reset := ansi.ColorCode("reset")
	for i, t := range tasks {
        fmt.Printf("%2d [ %s%10s%s ] %s@%s%s: ", i,red, t.Due_on, reset, green, t.Assignee.Name, reset)
        for _, p := range t.Projects {
            fmt.Printf("%s#%s%s ", blue,p.Name, reset)
        }

        for _, ta := range t.Tags {
            fmt.Printf("%s#%s%s ", blue, ta.Name, reset)
        }
        fmt.Printf("%s\n", t.Name)
	}
}

func cache(tasks []api.Task_t) {
	f, _ := os.Create(utils.CacheFile())
	defer f.Close()

    blue := ansi.ColorCode("blue")
    reset := ansi.ColorCode("reset")

	for i, t := range tasks {
		f.WriteString(strconv.Itoa(i) + ":")
		f.WriteString(strconv.Itoa(t.Id) + ":")
		f.WriteString(t.Due_on + ":")
        f.WriteString(t.Assignee.Name + ":")
        for _, p := range t.Projects {
            f.WriteString(blue + "#" + p.Name + " " + reset)
        }

        for _, ta := range t.Tags {
            f.WriteString(blue + "#" + ta.Name + " " + reset)
        }
		f.WriteString(t.Name + " ")
        f.WriteString("\n")
    }
}

func format(line string) {
    green := ansi.ColorCode("green")
    bold := ansi.ColorCode("white")
    red := ansi.ColorCode("red")
    reset := ansi.ColorCode("reset")
	dateRegexp := "[0-9]{4}-[0-9]{2}-[0-9]{2}"

	index := regexp.MustCompile("^[0-9]*").FindString(line)
	line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove index
	line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove task_id

    date := regexp.MustCompile("^" + dateRegexp).FindString(line)
	line = regexp.MustCompile("^("+dateRegexp+")?:").ReplaceAllString(line, "") // remove date

    assignee := regexp.MustCompile("^[a-zA-Z]*:").FindString(line)
	line = regexp.MustCompile("^.*:").ReplaceAllString(line, "") // remove assignee

	fmt.Printf("%s%2s%s [ %s%10s%s ] %s@%s%s %s\n", bold, index,reset, red, date, reset, green, assignee, reset, line)
}
