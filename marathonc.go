package main

import (
	"flag"
	"fmt"
	"github.com/antonlindstrom/marathonc/marathon"
)

func main() {
	var host = flag.String("host", "http://localhost:8080", "Marathon host")
	flag.Parse()

	// Get tasks
	tasks, err := marathon.GetTasks(*host)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// Print out tasks running
	for _, task := range tasks.List {
		fmt.Printf("%s => %s:%d\n", task.AppId, task.Host, task.Ports[0])
	}

	// Get apps
	apps, err := marathon.GetApps(*host)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// Print out apps
	for _, app := range apps.List {
		fmt.Printf("App: %s\n", app.Id)
		fmt.Printf("  Instances: %d\n", app.Instances)
		fmt.Printf("  Version: %s\n", app.Version)
		fmt.Printf("  Executor: %s\n", app.Executor)
		fmt.Printf("  Container:\n\tImage: %s\n\tOptions: %v\n", app.Container.Image, app.Container.Options)
		fmt.Printf("  Cpus: %f\n", app.Cpus)
		fmt.Printf("  Memory: %f\n", app.Mem)
	}
}
