package statscollector

import (
	"time"
)

type SystemStatsSummary struct {
	LastUpdated time.Time          `json:"last_updated"`
	Docker      DockerStatsSummary `json:"docker"`
}

// Main entrypoint for the collector.
//
// Calling this function will collect data from all the desired
// locations and produce a struct containing all the data.
//
// See SystemStatsSummary for information on the data format.
//
// Any error encountered is returned as well.
func CollectStatistics() (SystemStatsSummary, error) {
	dockerStats, dockerErr := collectDockerStats()

	if dockerErr != nil {
		return SystemStatsSummary{}, dockerErr
	}

	return SystemStatsSummary{Docker: dockerStats, LastUpdated: time.Now()}, nil

}
