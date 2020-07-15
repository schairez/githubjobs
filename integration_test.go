package githubjobs

import (
	"testing"
)

func TestGetPositionsResultStructFromAPI(t *testing.T) {
	lat := "37.3229978"
	long := "-122.0321823"
	var params = OptParams{Lat: &lat, Long: &long}
	c := NewClient()
	c.SetGithubJobsURLFromStruct(&params)
	jobsData, err := c.GetPositionsResultStruct()
	if err != nil {
		t.Errorf("Unexpected error for empty API request result, jobsData=%v", jobsData)
	}

}

//below test ensures fn fails when no return from API b/c of no jobs given params
func TestBuildQueryFromMap_AndReturnEmptyAPIResultError(t *testing.T) {
	paramsMap := map[string]string{
		"description": "restroom attendant",
		"location":    "Oaxaca",
		"full_time":   "true"}
	c := NewClient()
	c.SetGithubJobsURLFromMap(paramsMap)
	jobsData, err := c.GetPositionsResultStruct()
	if err == nil {
		t.Errorf("Expected error for empty API request result, jobsData=%v", jobsData)
	}

}
