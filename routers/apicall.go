package routers

import (
	"github.com/basit9958/shell-command-executor/models"
	"github.com/gofiber/fiber/v2"
	"os/exec"
)

func Execute(app *fiber.Ctx) error {
	requestBody := new(models.Request)

	if err := app.BodyParser(&requestBody); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot Parse the Request ",
		})
	}
	command := app.Query("command")
	if command == "" {
		command = requestBody.Command
	}

	// Execute the command
	output, err := executeCommand(command)
	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		if ok && exitError.ExitCode() == 127 {
			// Command not found
			return fiber.NewError(fiber.StatusNotFound, "Command not found")
		}
		// Other error occurred
		return fiber.NewError(fiber.StatusInternalServerError, "Command execution failed")
	}

	// Create the response
	response := models.Response{
		Output: string(output),
	}

	// Return the response as JSON
	return app.JSON(response)
}
