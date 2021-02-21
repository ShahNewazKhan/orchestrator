package models

import (
	"errors"
	"log"

	"github.com/Kamva/mgm/v2"
)

type JobStatus string

const (
	Pending = "PENDING"
	Started = "STARTED"
	Running = "RUNNING"
	Errored = "ERRORED"
	Done    = "DONE"
)

type Job struct {
	mgm.DefaultModel `bson:",inline"`
	Status           string `json:"status" bson:"status"`
	Name             string `json:"name" bson:"name"`
	Completed        bool   `json:"complete" bson:"completed"`
	BuildId          string `json:"buildId" bson:"buildId"`
	WorkerId         string `json:"workerId" bson:"workerId"`
}

func CreateJob(name string) *Job {
	return &Job{
		Status:    "PENDING",
		Name:      name,
		Completed: false,
	}
}

func (js JobStatus) IsValid() error {
	//TODO: check jobstatus against a state machine based on current status
	switch js {

	case Pending, Started, Running, Errored, Done:
		log.Printf("%s is valid", js)
		return nil
	}

	return errors.New("Invalid job status")
}
