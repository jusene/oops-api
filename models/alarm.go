package models

// alertmanager 告警结构体

type AlertManager struct {
	TITLE string `json:"title"`
	INSTANCE string ``
	STATUS string
	SEVERITY string
	JOB string
	TIME string
}