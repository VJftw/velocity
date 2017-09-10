package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/velocity-ci/velocity/master/velocity/domain"
	"github.com/velocity-ci/velocity/master/velocity/domain/task"
)

func main() {
	version := flag.Bool("v", false, "Show version")
	list := flag.Bool("l", false, "List tasks")

	flag.Parse()

	if *version {
		fmt.Println("Version")
		os.Exit(0)
	} else if *list {
		// look for task ymls and parse them into memory.
		tasks := getTasksFromDirectory("./tasks/")
		// iterate through tasks in memory and list them.
		for _, task := range tasks {
			fmt.Printf("%s: %s (", task.Name, task.Description)
			for _, parameter := range task.Parameters {
				fmt.Printf(" %s= %s ", parameter.Name, parameter.Value)
			}
			fmt.Println(")")
			for _, step := range task.Steps {
				fmt.Printf("\t%s| %s: %s\n", step.GetType(), step.GetDescription(), step.GetDetails())
			}
		}
		os.Exit(0)
	}

	switch os.Args[1] {
	case "run":
		run(os.Args[2])
		break
	}
}

func getTasksFromDirectory(dir string) []domain.Task {
	tasks := []domain.Task{}

	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".yml") || strings.HasSuffix(f.Name(), ".yaml") {
			taskYml, _ := ioutil.ReadFile(fmt.Sprintf("%s%s", dir, f.Name()))
			task := task.ResolveTaskFromYAML(string(taskYml))
			tasks = append(tasks, task)
		}
		return nil
	})

	return tasks
}

func run(taskName string) {
	tasks := getTasksFromDirectory("./tasks/")

	var task *domain.Task
	// find Task requested
	for _, t := range tasks {
		if t.Name == taskName {
			task = &t
			break
		}
	}

	if task == nil {
		panic(fmt.Sprintf("Task %s not found\n%s", taskName, tasks))
	}

	fmt.Printf("Running task: %s (from: %s)\n", task.Name, taskName)

	// Resolve parameters
	reader := bufio.NewReader(os.Stdin)
	resolvedParams := []domain.Parameter{}
	for _, p := range task.Parameters {
		// get real value for parameter (ask or from env)
		inputText := ""
		for len(strings.TrimSpace(inputText)) < 1 {
			fmt.Printf("Enter a value for %s (default: %s): ", p.Name, p.Value)
			inputText, _ = reader.ReadString('\n')
		}
		p.Value = strings.TrimSpace(inputText)
		resolvedParams = append(resolvedParams, p)
	}
	task.Parameters = resolvedParams
	task.UpdateParams()
	task.SetEmitter(func(s string) { fmt.Printf("    %s\n", s) })

	// Run each step unless they fail (optional)
	for _, step := range task.Steps {
		step.Execute()
	}
}
