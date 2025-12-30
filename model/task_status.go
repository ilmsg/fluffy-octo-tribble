package model

type TaskStatus string

const (
	Process TaskStatus = "Process"
	Pending TaskStatus = "Pending"
	Success TaskStatus = "Success"
)
