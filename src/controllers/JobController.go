package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"orchestrator/models"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllJobs - GET /api/jobs
func GetAllJobs(ctx *fiber.Ctx) {
	collection := mgm.Coll(&models.Job{})
	jobs := []models.Job{}

	err := collection.SimpleFind(&jobs, bson.D{})

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})

		return // necessary, or else controller will continue
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"jobs": jobs,
	})
}

// GetJobByID - GET /api/jobs/:id
func GetJobByID(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	job := &models.Job{}
	collection := mgm.Coll(job)

	err := collection.FindByID(id, job)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": fmt.Sprintf("Job %s not found.", id),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":  true,
		"job": job,
	})
}

// CreateJob - POST /api/jobs
func CreateJob(ctx *fiber.Ctx) {
	params := new(struct {
		BrigadeProject string
		BrigadeSecret  string
		Name           string
		VideoUrl	   string
	})

	ctx.BodyParser(&params)

	if len(params.BrigadeProject) == 0 || len(params.Name) == 0 || len(params.BrigadeSecret) == 0 || len(params.VideoUrl) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Project, secret, url or name not specified.",
		})
		return
	}

	job := models.CreateJob(params.Name, params.VideoUrl)
	err := mgm.Coll(job).Create(job)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	go triggerBuild(job.IDField.ID.Hex(), params.BrigadeProject, params.BrigadeSecret, params.VideoUrl)

	ctx.JSON(fiber.Map{
		"ok":  true,
		"job": job,
	})

}

func triggerBuild(jobId, brigadeProject, brigadeSecret, videoUrl string) {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"jobId": jobId,
		"videoUrl" : videoUrl,
	})
	responseBody := bytes.NewBuffer(postBody)

	baseBrigadeUrl := os.ExpandEnv("http://$GENERIC_GATEWAY_HOST:$GENERIC_GATEWAY_PORT/simpleevents/v1")
	brigadeUrl := fmt.Sprintf("%s/%s/%s", baseBrigadeUrl, brigadeProject, brigadeSecret)
	log.Printf("Calling brigade at %v", brigadeUrl)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(brigadeUrl, "application/json", responseBody)

	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

// ToggleJobStatus - PATCH /api/jobs/:id
func ToggleJobStatus(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	job := &models.Job{}
	collection := mgm.Coll(job)

	err := collection.FindByID(id, job)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Job not found.",
		})
		return
	}

	job.Completed = !job.Completed

	err = collection.Update(job)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":  true,
		"job": job,
	})
}

// UpdateJobStatus - PATCH /api/jobs/:id/status
func UpdateJobStatus(ctx *fiber.Ctx) {
	params := new(struct {
		Status string
	})

	ctx.BodyParser(&params)

	jobStatus := models.JobStatus(params.Status)
	if err := jobStatus.IsValid(); err != nil {
		ctx.Status(422).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid status, valid statuses ['PENDING', 'STARTED', 'RUNNING','ERRORED','DONE']",
		})
		return
	}

	id := ctx.Params("id")

	job := &models.Job{}
	collection := mgm.Coll(job)

	err := collection.FindByID(id, job)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Job not found.",
		})
		return
	}

	job.Status = params.Status

	err = collection.Update(job)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":  true,
		"job": job,
	})
}

// UpdateJobDetails - PATCH /api/jobs/:id/brigade
func UpdateJobDetails(ctx *fiber.Ctx) {
	params := new(struct {
		BuildId  string
		WorkerId string
	})

	ctx.BodyParser(&params)

	id := ctx.Params("id")

	job := &models.Job{}
	collection := mgm.Coll(job)

	err := collection.FindByID(id, job)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Job not found.",
		})
		return
	}

	job.BuildId = params.BuildId
	job.WorkerId = params.WorkerId

	err = collection.Update(job)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":  true,
		"job": job,
	})
}

// DeleteJob - DELETE /api/jobs/:id
func DeleteJob(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	job := &models.Job{}
	collection := mgm.Coll(job)

	err := collection.FindByID(id, job)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Job not found.",
		})
		return
	}

	err = collection.Delete(job)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":  true,
		"job": job,
	})
}
