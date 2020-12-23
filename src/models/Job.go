package models

import (
    "github.com/Kamva/mgm/v2"
)

type Job struct {
	mgm.DefaultModel `bson:",inline"`
	Status string `json:"status" bson:"status"`
	Type string `json:"type" bson:"type"` 
	Completed bool `json:"complete" bson:"completed"`
}

func CreateJob(status, type: string) *Job {
	return &Job(
		Status: status,
		Type: type, 
		Completed: false,
	)
}