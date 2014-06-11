package marathon

import (
	"encoding/json"
)

// Marathon tasks
type Tasks struct {
	List []Task `json:"tasks"`
}

// A marathon task
type Task struct {
	AppId   string `json:"appId"`
	Id      string `json:"id"`
	Host    string `json:"host"`
	Ports   []int  `json:"ports"`
	Started string `json:"startedAt"`
	Staged  string `json:"stagedAt"`
	Version string `json:"version"`
}

// Get tasks from endpoint
func GetTasks(host string) (*Tasks, error) {
	body, err := GetBody(host + "/v2/tasks")
	if err != nil {
		return nil, err
	}

	// Unmarshal
	var tasks Tasks
	err = json.Unmarshal(body, &tasks)

	return &tasks, err
}

// Get tasks for app
func GetAppTasks(host, id string) (*Tasks, error) {
	body, err := GetBody(host + "/v2/apps/" + id + "/tasks")
	if err != nil {
		return nil, err
	}

	// Unmarshal
	var tasks Tasks
	err = json.Unmarshal(body, &tasks)

	return &tasks, err
}
