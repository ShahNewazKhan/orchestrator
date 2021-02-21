package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gofiber/fiber"
)

type Projects []struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	GenericGatewaySecret string `json:"genericGatewaySecret"`
}

// GetAllProjects - GET /api/projects
func GetAllProjects(ctx *fiber.Ctx) {
	url := os.ExpandEnv("http://$BRIGADE_API_HOST:$BRIGADE_API_PORT/v1/projects")

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	res, err := client.Do(req)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	var p Projects
	err = json.NewDecoder(res.Body).Decode(&p)

	ctx.JSON(fiber.Map{
		"ok":       true,
		"projects": p,
	})
}
