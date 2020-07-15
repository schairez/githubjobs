package githubjobs

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

/*

description — A search term, such as "ruby" or "java". This parameter is aliased to search.
location — A city name, zip code, or other location search term.
lat — A specific latitude. If used, you must also send long and must not send location.
long — A specific longitude. If used, you must also send lat and must not send location.
full_time — If you want to limit results to full time positions set this parameter to 'true'.

*/

func buildGithubJobsURLFromMap(uriParams map[string]string) string {
	u, err := url.Parse(BaseURLGithubPositionsAPI)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	if uriParams != nil {
		for k, v := range uriParams {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}
	return u.String()
}

func buildGithubJobsURLFromStruct(params *OptParams) string {

	u, err := url.Parse(BaseURLGithubPositionsAPI)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	if params != nil {
		if params.Description != nil {
			q.Add("description", *params.Description)
		}
		if params.Location != nil {
			q.Add("location", *params.Location)
		}
		if params.FullTime != nil {
			q.Set("full_time", *params.FullTime)
		}
		if params.FullTime != nil {
			q.Set("full_time", *params.FullTime)

		}
		if params.Lat != nil {
			q.Set("lat", *params.Lat)

		}
		if params.Long != nil {
			q.Set("long", *params.Long)

		}

		u.RawQuery = q.Encode()
	}
	return u.String()

}

//NewClient sets up a new http Client for the Github Jobs API
func NewClient() *Client {
	return &Client{
		Client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

//SetGithubJobsURLFromMap receives a list of objs representing the params that
// will be passed to Github Jobs REST API
func (c *Client) SetGithubJobsURLFromMap(uriParams map[string]string) {
	c.JobsURLApi = buildGithubJobsURLFromMap(uriParams)

}

//SetGithubJobsURLFromStruct receives a list of objs representing the params that
// will be passed to Github Jobs REST API
func (c *Client) SetGithubJobsURLFromStruct(params *OptParams) {
	c.JobsURLApi = buildGithubJobsURLFromStruct(params)

}

//SetHTTPClient sets a new http client
func (c *Client) SetHTTPClient(client *http.Client) {
	c.Client = client
}

/*
GetPositionsResultStruct sends a GET request to the Github Jobs API and returns the json response as a ArrOfJobListingStructs
GET /positions.json
*/
func (c *Client) GetPositionsResultStruct() (*ArrOfJobListingStructs, error) {
	target := &ArrOfJobListingStructs{}
	resp, err := c.Client.Get(c.JobsURLApi)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	//a non-2xxx response doesn't cause an error w/ http client
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		return nil, fmt.Errorf("non-2xxx with status_code=%v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, err
	}

	if len(*target) == 0 {
		return nil, errors.New("Empty structure b/c no results found for your query entry")

	}

	return target, nil
}
