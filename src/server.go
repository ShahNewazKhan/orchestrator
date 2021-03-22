package main

import (
	"log"
	"os"

	"orchestrator/controllers"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

	connectionString := os.ExpandEnv("mongodb://$MONGODB_USER:$MONGODB_PASS@$MONGODB_HOST:$MONGODB_PORT")
	log.Printf("mongo connection string: %s", connectionString)

	err := mgm.SetDefaultConfig(nil, "jobs", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger(middleware.LoggerConfig{
		Format: "${time} ${method} ${requestid} ${path}",
	}))

	// jobs api
	app.Get("/api/jobs", controllers.GetAllJobs)
	app.Get("/api/jobs/:id", controllers.GetJobByID)
	app.Post("/api/jobs", controllers.CreateJob)
	app.Patch("/api/jobs/:id", controllers.ToggleJobStatus)
	app.Patch("/api/jobs/:id/brigade", controllers.UpdateJobDetails)
	app.Patch("/api/jobs/:id/status", controllers.UpdateJobStatus)
	app.Delete("/api/jobs/:id", controllers.DeleteJob)

	// projects api
	app.Get("/api/projects", controllers.GetAllProjects)
	log.Fatal(app.Listen(3000))
	
}
