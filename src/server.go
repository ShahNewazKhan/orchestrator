package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"github.com/jozsefsallai/fiber-todo-demo/controllers"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

	connectionString := os.ExpandEnv("mongodb://$MONGODB_USER:$MONGODB_PASS@MONGODB_HOST:$MONGODB_PORT")
	log.Printf("mongo connection string: %s", connectionString)

	err := mgm.SetDefaultConfig(nil, "jobs", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	app.Get("/api/jobs", controllers.GetAllTodos)
	app.Get("/api/jobs/:id", controllers.GetTodoByID)
	app.Post("/api/jobs", controllers.CreateTodo)
	app.Patch("/api/jobs/:id", controllers.ToggleTodoStatus)
	app.Delete("/api/jobs/:id", controllers.DeleteTodo)

	app.Listen(3000)
}
