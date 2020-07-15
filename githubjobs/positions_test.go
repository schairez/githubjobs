package githubjobs

import (
	"net/http"
	"testing"
	"time"
)

func TestBuildGithubJobsURLFromMap(t *testing.T) {
	paramsMap := map[string]string{"description": "java", "location": "sf", "full_time": "true"}
	expected := "https://jobs.github.com/positions.json?description=java&full_time=true&location=sf"
	actual := buildGithubJobsURLFromMap(paramsMap)
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}

}

func TestBuildGithubJobsURLFromStruct(t *testing.T) {
	lat := "37.3229978"
	long := "-122.0321823"
	expected := "https://jobs.github.com/positions.json?lat=37.3229978&long=-122.0321823"
	var params = OptParams{Lat: &lat, Long: &long}
	actual := buildGithubJobsURLFromStruct(&params)
	t.Log(actual)
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}

}

//below test ensures fn fails safely when no return from API
func TestBuildQueryFromMap_EmptyAPIResultError(t *testing.T) {
	paramsMap := map[string]string{
		"description": "restroom attendant",
		"location":    "jupyter",
		"full_time":   "true"}

	var c = &JobsClient{
		Client: &http.Client{Timeout: 5 * time.Second}}
	c.SetGithubJobsURLFromMap(paramsMap)
	jobsData, err := c.GetPositionsResultStruct()
	if err == nil {
		t.Errorf("Expected error for empty API request, jobsData=%v", jobsData)
	}
	if err != nil {
		t.Log(err)

	}

}

// func TestBuildQuery(t *testing.T) {
// 	req, err := http.NewRequest("GET", BaseURLGithubPositionsAPI, nil)
// 	if err != nil {
// 		log.Print(err)
// 		os.Exit(1)
// 	}
// 	q := url.Values{}
// 	paramsMap := map[string]string{"description": "python is great", "location": "sf", "full_time": "true"}
// 	for k, v := range paramsMap {
// 		q.Add(k, v)
// 	}

// 	t.Log(q)
// 	req.URL.RawQuery = q.Encode()

// 	t.Log(req.URL.String())
// 	var c = &JobsClient{
// 		Client: &http.Client{Timeout: 5 * time.Second}}
// 	resp, err := c.Client.Do(req)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	defer resp.Body.Close()
// 	var target []JobsListing
// 	decoder := json.NewDecoder(resp.Body)
// 	err = decoder.Decode(&target)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// t.Log(target)

// }

// func TestGetJSONresponse(t *testing.T) {

// 	paramsMap := map[string]string{
// 		"description": "python",
// 		"location":    "sf",
// 		"full_time":   "true"}

// 	var c = &JobsClient{
// 		Client: &http.Client{Timeout: 5 * time.Second}}
// 	c.SetGithubJobsURL(paramsMap)
// 	var jobsData []JobsListing
// 	jobsData = c.GetJSONResponse()
// 	// t.Log(jobsData)
// 	t.Log(jobsData[0].Description)
// 	t.Log(stripHTML(jobsData[0].Description))
// 	// t.Log(c.GetJSONResponse())

// }

// func TestStringHTMLSplitter(t *testing.T) {

// 	var body = `
//   <p><strong>WORKING AT REBILLY</strong></p> <p>Rebilly’s Purpose: Find and create freedom through your work.</p> <p>Rebilly offers a competitive salary, training and development, birthday lunches, and provides the computer of your choosing. Further perks and benefits are dependent on your work location of choice.</p> <p>Rebilly is a Teal Organization, meaning our organizational structure may differ than what you’re used to, but we think you’ll be pleasantly surprised. In a very tiny nutshell, this means that as an organization we value self-management, self-organization, as well as the wholeness of the individuals that make up our team (meaning you should be yourself at work, and do the work that inspires you.)</p> <p>Come check us out at <a href="https://www.rebilly.com/careers/">https://www.rebilly.com/careers/</a></p> <p><strong>ABOUT REBILLY</strong></p> <p>Rebilly is now a payments software after pivoting from a focus on subscription billing in October 2019. Rebilly’s comprehensive feature set is built to help our merchants get more of their customers from the order form to the thank you page, and more renewals paid. We took lessons learned from years of experience in the subscription business and millions of subscribers and applied them to make our payment system clever, flexible, and easy to use for our clients.</p> ",
// how_to_apply: "<p>Please send your resume through Apply By API: <a href="https://app.applybyapi.com/posting/7/">https://app.applybyapi.com/posting/7/</a></p>
// 	 `
// 	t.Log(stripHTML(body))

// }
