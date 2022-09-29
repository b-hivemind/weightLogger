package api

type entriesQuery struct {
	Days int `binding:"required" uri:"numdays"`
}

type newEntryQuery struct {
	Weight float32 `json:"weight" binding:"required"`
	Force  bool    `json:"force"`
}
