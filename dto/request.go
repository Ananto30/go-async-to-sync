package dto

// DemoInfoReq is sample request to the go server
type DemoInfoReq struct {
	ArticleID string `json:"articleId"`
}

// DemoAsyncReq is sample async request to the pyhton server
type DemoAsyncReq struct {
	TrackID string `json:"trackId"`
}