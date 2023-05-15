package statscollector

import (
	"time"
)

type SystemStatsSummary struct {
	LastUpdated time.Time          `json:"last_updated"`
	Docker      DockerStatsSummary `json:"docker"`
}

func CollectStatistics() (SystemStatsSummary, error) {
	dockerStats, dockerErr := collectDockerStats()

	if dockerErr != nil {
		return SystemStatsSummary{}, dockerErr
	}

	return SystemStatsSummary{Docker: dockerStats, LastUpdated: time.Now()}, nil

}
