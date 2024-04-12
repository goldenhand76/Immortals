package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Notification struct {
	From    User   `json:"from"`
	To      User   `json:"to"`
	Message string `json:"message"`
}

type NodeData struct {
	NodeID   string            `json:"nodeId,omitempty"` // Node ID that set on device and not changable Use omitempty to omit this field if it's empty
	Sensor   map[string]string `json:"sensor"`           // List of sensors that exists and their topics
	Actuator map[string]string `json:"actuator"`         // List of actuators that exists on node and their topics
}

type Device struct {
	ID    int
	Name  string
	Value int
}

type DetectedDevice struct {
	ID        int
	Topic     string
	Type      string
	is_online bool
}
