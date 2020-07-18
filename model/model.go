package model

type GenerateRequest struct {
	Text string `json:"text"`
	Size int    `json:"size"`
}

type GenerateResponse struct {
	Url string `json:"url,omitempty"`
	Err string `json:"error,omitempty"`
}
