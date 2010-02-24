package makengo

import (
	"fmt"
	"os"
	"flag"
)

type runner struct {
	usage func()
}

func NewRunner() *runner {
	custom_usage := func() {
		fmt.Fprintf(os.Stderr, "makengo [options] targets...\n\n")
		fmt.Fprintf(os.Stderr, "Options are:\n\n")
		flag.PrintDefaults()
	}

	return &runner{usage: custom_usage}
}

func (self *runner) Run() {
	var help *bool = flag.Bool("h", false, "Show usage")
	var showDescriptions *bool = flag.Bool("T", false, "Show task descriptions")

	flag.Usage = self.usage
	flag.Parse()

	if *help == true {
		flag.Usage()
		return
	}

	if *showDescriptions == true {

		for _, task := range TaskManager {
			if task.Description != "" {
				fmt.Printf("makengo %-10s # %s\n", task.Name, task.Description)
			}
		}

		return
	}

	if len(flag.Args()) > 0 {
		TaskManager.InvokeByName(flag.Args())
	} else {
		TaskManager.InvokeByName([]string{"Default"})
	}
}
