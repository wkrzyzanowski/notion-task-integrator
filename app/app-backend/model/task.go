package model

type Task struct {
	Id          string
	StartDate   string
	EndDate     string
	Interval    string
	Name        string
	Description string
	Metadata    map[string]string
}
