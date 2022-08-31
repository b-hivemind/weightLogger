package api

type entriesQuery struct {
	days int `binding:"required" uri:"numdays"`
}

type Response_New struct {
	Weight string `json:"weight"`
	Force  bool   `json:"force"`
}
