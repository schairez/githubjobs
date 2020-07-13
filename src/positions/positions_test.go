package positions

import (
	"strings"
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

// func TestGetJSONresponse(t *testing.T) {
// 	// jobsURL := "https://jobs.github.com/positions.json?description=python&location=sf&full_time=true"

// 	paramsMap := map[string]string{
// 		"description": "python",
// 		"location":    "sf",
// 		"full_time":   "true"}

// 	var c = &JobsClient{
// 		Client: &http.Client{Timeout: 5 * time.Second}}
// 	c.SetGithubJobsURL(paramsMap)
// 	t.Log(c.GetJSONResponse())

// }

func TestStringHTMLSplitter(t *testing.T) {

	str := "<p>The Job:</p> <p>Traction is looking for a Lead Full Stack Developer. You should have prior experience architecting and developing advanced dynamic projects involving technologies for Front-end, Back-end, Database, DevOps and Mobile, as well as leading teams of engineers working on these projects.</p> <p>This role will see you leading a team of developers as well as being an individual contributor to the team.</p> <p>We’re on a mission to become one of the most in-demand agencies in the United States, sought out by the best brands in the world to design meaningful experiences for their customers.</p> <p>Must be available to work onsite in Sunnyvale, CA exclusively working embedded with our client that is headquartered in Cupertino.</p> <p>Core Responsibilities:</p> <ul> <li>Act as the primary technical contact for an engineering team with other teams across disciplines, potentially leading a team of 2-5 developers</li> <li>Work closely with the development manager to ensure requirements and timelines are met</li> <li>Work closely with the development manager and back-end teams to architect and integrate front-end code</li><li>Knowledge of common design patterns in web development (i.e., MVC, MVVM)</li> <li>Working knowledge of source control software such as Git (preferred) or SVN</li> <li>Exposure to Agile project management methodologies</li> <li>Ability to multi-task and manage tasks with varying priorities</li></ul> <p>Our humblebrag: Join an award-winning, digitally-focused, full-service, independent agency staffed by a bunch of folks who love what they do and do what they love. Still on the rise after 17 years of success, you’ll enjoy a satisfying suite of benefits, as well as the chance to work on the world's most well-know brand.</p>"

	split := strings.Split(str, "<p>")
	t.Log(split)

}
