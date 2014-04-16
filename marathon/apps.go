package marathon

import (
	"encoding/json"
)

// Marathon tasks
type Apps struct {
	List []App `json:"apps"`
}

// A marathon app
type App struct {
	Cmd           string            `json:"cmd"`
	Constraints   [][]string        `json:"constaints"`
	Container     Container         `json:"container"`
	Cpus          float32           `json:"cpus"`
	Env           map[string]string `json:"env"`
	Executor      string            `json:"executor"`
	Id            string            `json:"id"`
	Instances     int               `json:"instances"`
	Mem           float32           `json:"mem"`
	Ports         []int             `json:"ports"`
	TaskRateLimit float32           `json:"tasksRateLimit"`
	TasksRunning  int               `json:"tasksRunning"`
	TasksStaged   int               `json:"tasksStaged"`
	Uris          []string          `json:"uris"`
	Version       string            `json:"version"`
}

// A marathon app container
type Container struct {
	Image   string   `json:"image"`
	Options []string `json:"options"`
}

// Get apps from endpoint
func GetApps(host string) (*Apps, error) {
	body, err := GetBody(host + "/v2/apps")
	if err != nil {
		return nil, err
	}

	// Unmarshal
	var apps Apps
	err = json.Unmarshal(body, &apps)

	return &apps, err
}
