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
	Cmd           string            `json:"cmd,omitempty"`
	Constraints   [][]string        `json:"constaints,omitempty"`
	Container     *Container        `json:"container,omitempty"`
	Cpus          float32           `json:"cpus,omitempty"`
	Env           map[string]string `json:"env,omitempty"`
	Executor      string            `json:"executor,omitempty"`
	Id            string            `json:"id,omitempty"`
	Instances     int               `json:"instances,omitempty"`
	Mem           float32           `json:"mem,omitempty"`
	Ports         []int             `json:"ports,omitempty"`
	TaskRateLimit float32           `json:"tasksRateLimit,omitempty"`
	TasksRunning  int               `json:"tasksRunning,omitempty"`
	TasksStaged   int               `json:"tasksStaged,omitempty"`
	Uris          [0]string         `json:"uris"`
	Version       string            `json:"version,omitempty"`
}

// A marathon app container
type Container struct {
	Image   string    `json:"image"`
	Options [0]string `json:"options"`
}

// Message, result from post
type Msg struct {
	Message string `json:"message"`
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

// Create new app at endpoint
func NewApp(host string, app *App) (*Msg, error) {
	// Unmarshal
	b, err := json.Marshal(&app)
	if err != nil {
		return nil, err
	}

	// Send request
	body, err := PostBody(host+"/v2/apps", b)
	if err != nil {
		return nil, err
	}

	// Unmarshal if body exists, body is empty if response
	// is ok.
	if string(body) != "" {
		var msg Msg
		err = json.Unmarshal(body, &msg)
		return &msg, err
	}

	return &Msg{Message: "OK"}, nil
}

// Delete an app
func DestroyApp(host, id string) (*Msg, error) {
	// Send request
	body, err := DeleteBody(host+"/v2/apps/"+id, []byte(""))
	if err != nil {
		return nil, err
	}

	// Unmarshal if body exists, body is empty if response
	// is ok.
	if string(body) != "" {
		var msg Msg
		err = json.Unmarshal(body, &msg)
		return &msg, err
	}

	return &Msg{Message: "OK"}, nil
}
