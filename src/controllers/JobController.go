package controllers

import (
    "github.com/Kamva/mgm/v2"
    "github.com/gofiber/fiber"
    "orchestrator/models"
    "go.mongodb.org/mongo-driver/bson"
    "fmt"
)

// GetAllJobs - GET /api/jobs
func GetAllJobs(ctx *fiber.Ctx) {
    collection := mgm.Coll(&models.Job{})
    jobs = []models.Job{}

    err := collection.SimpleFind(&jobs, bson.D{})

    if err != nil {
        ctx.Status(500).JSON(fibre.Map{
            "ok" : false,
            "error" : err.Error(),
        })

        return // necessary, or else controller will continue
    }
    
    ctx.JSON(fibre.Map{
        "ok": true,
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
        "ok":   true,
        "job": job,
    })
}


// CreateJob - POST /api/jobs
func CreateJob(ctx *fiber.Ctx) {
    params := new(struct {
        Status string
        Type string
    })

    ctx.BodyParser(&params)

    if len(params.Title) == 0 || len(params.Description) == 0 {
        ctx.Status(400).JSON(fiber.Map{
            "ok":    false,
            "error": "Status or type not specified.",
        })
        return
    }

    job := models.CreateJob(params.Status, params.Type)
    err := mgm.Coll(job).Create(job)
    if err != nil {
        ctx.Status(500).JSON(fiber.Map{
            "ok":    false,
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(fiber.Map{
        "ok":   true,
        "job": job,
    })
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

    job.Done = !job.Done

    err = collection.Update(job)
    if err != nil {
        ctx.Status(500).JSON(fiber.Map{
            "ok":    false,
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(fiber.Map{
        "ok":   true,
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
        "ok":   true,
        "job": job,
    })
}