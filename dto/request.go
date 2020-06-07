package dto

// DemoBalanceReq is sample request to the go server
type DemoBalanceReq struct {
	AccountID string `json:"accountId"`
}

// DemoAsyncReq is sample async request to the pyhton server
type DemoAsyncReq struct {
	TrackID   string `json:"trackId"`
	AccountID string `json:"accountId"`
}
