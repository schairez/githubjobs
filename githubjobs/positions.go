package githubjobs

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	//BaseURLGithubPositionsAPI comment
	BaseURLGithubPositionsAPI = "https://jobs.github.com/positions.json?"
)

func buildGithubJobsURL(uriParams map[string]string) string {
	size := len(uriParams)
	cnt := 0
	urlparams := ""
	for k, v := range uriParams {
		urlparams += k + "=" + v
		cnt++
		if size != cnt {
			urlparams += "&"
		}

	}
	return BaseURLGithubPositionsAPI + urlparams

}

//SetHTTPClient sets a new http client
func (c *JobsClient) SetHTTPClient(client *http.Client) {
	c.Client = client
}

//SetGithubJobsURL receives a list of objs representing the params that
// will be passed to Github Jobs REST API
func (c *JobsClient) SetGithubJobsURL(uriParams map[string]string) {
	c.JobsURLApi = buildGithubJobsURL(uriParams)

}

/*

description — A search term, such as "ruby" or "java". This parameter is aliased to search.
location — A city name, zip code, or other location search term.
lat — A specific latitude. If used, you must also send long and must not send location.
long — A specific longitude. If used, you must also send lat and must not send location.
full_time — If you want to limit results to full time positions set this parameter to 'true'.

*/

//GET /positions.json

//GetJSONResponse returns structs of the json GET REST API
func (c *JobsClient) GetJSONResponse() []JobsListing {
	var target []JobsListing
	resp, err := c.Client.Get(c.JobsURLApi)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("response status code was %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&target)
	if err != nil {
		log.Fatalln(err)
	}
	return target
}
