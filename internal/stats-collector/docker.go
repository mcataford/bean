package statscollector

import (
	"bytes"
	"encoding/json"
	"io"
	"os/exec"
)

type DockerPsOutput struct {
	Names      string `json:"Names"`
	State      string `json:"State"`
	Ports      string `json:"Ports"`
	RunningFor string `json:"RunningFor"`
}

type DockerServiceStats struct {
	Name       string `json:"name"`
	RunningFor string `json:"uptime"`
	State      string `json:"state"`
}

type DockerStatsSummary struct {
	Services map[string]DockerServiceStats `json:"services"`
}

// Collects and formats data from `docker ps`. This requires `jq` to be installed
// since the output of `docker ps` requires slurping to form a valid json array.
//
// This may return a report or an error, depending on circumstance.
func collectDockerStats() (DockerStatsSummary, error) {
	dockerPs := exec.Command("docker", "ps", "--format", "json")
	jqSlurp := exec.Command("jq", "-s")

	read, write := io.Pipe()
	dockerPs.Stdout = write
	jqSlurp.Stdin = read

	var outputBuffer bytes.Buffer
	jqSlurp.Stdout = &outputBuffer

	dockerPs.Start()
	jqSlurp.Start()
	dockerPs.Wait()
	write.Close()
	jqSlurp.Wait()

	var formattedOutput []DockerPsOutput
	jsonErr := json.Unmarshal(outputBuffer.Bytes(), &formattedOutput)

	if jsonErr != nil {
		return DockerStatsSummary{}, jsonErr
	}

	serviceStats := DockerStatsSummary{Services: map[string]DockerServiceStats{}}
	for _, dockerService := range formattedOutput {
		serviceStats.Services[dockerService.Names] = DockerServiceStats{Name: dockerService.Names, RunningFor: dockerService.RunningFor, State: dockerService.State}
	}

	return serviceStats, nil
}
