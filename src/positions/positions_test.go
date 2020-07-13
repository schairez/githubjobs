package positions

import (
	"testing"
)

func TestBuildGithubJobsURL(t *testing.T) {
	paramsMap := map[string]string{"description": "python", "location": "sf", "full_time": "true"}
	jobsURL := buildGithubJobsURL(paramsMap)
	t.Log(jobsURL)
	if jobsURL != "https://jobs.github.com/positions.json?description=python&location=sf&full_time=true" {
		t.Error()
	}

}
