package positions

import (
	"net/http"
	"testing"
	"time"
)

func TestBuildGithubJobsURL(t *testing.T) {
	paramsMap := map[string]string{"description": "python", "location": "sf", "full_time": "true"}
	jobsURL := buildGithubJobsURL(paramsMap)
	t.Log(jobsURL)
	if jobsURL != "https://jobs.github.com/positions.json?description=python&location=sf&full_time=true" {
		t.Error()
	}

}

func TestGetJSONresponse(t *testing.T) {
	// jobsURL := "https://jobs.github.com/positions.json?description=python&location=sf&full_time=true"

	paramsMap := map[string]string{
		"description": "python",
		"location":    "sf",
		"full_time":   "true"}

	var c = &JobsClient{
		Client: &http.Client{Timeout: 5 * time.Second}}
	c.SetGithubJobsURL(paramsMap)
	t.Log(c.GetJSONResponse())

}
