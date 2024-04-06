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
