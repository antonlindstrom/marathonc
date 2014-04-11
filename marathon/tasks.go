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
