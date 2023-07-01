package routers

import (
	"os/exec"
	"strings"
)

func executeCommand(cmd string) ([]byte, error) {
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}
	files := strings.Split(string(output), "\n")

	// Remove empty strings (if any)
	var validFiles []string
	for _, file := range files {
		if file != "" {
			validFiles = append(validFiles, file)
		}
	}

	output = []byte(strings.Join(validFiles, ", "))

	return output, nil
}
