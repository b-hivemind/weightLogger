package api

type entriesQuery struct {
	Days int `binding:"required" uri:"numdays"`
}

type newEntryQuery struct {
	Weight float32 `binding:"required" json:"weight"`
	Force  bool    `json:"force"`
}

type authQuery struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}
