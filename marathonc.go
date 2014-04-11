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
}
