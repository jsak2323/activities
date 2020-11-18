package rest

type (
	RequestActivity struct {
		ID     int64  `json:"id"`
		Name   string `json:"name"`
		Act    string `json:"act"`
		Status int    `json:"status"`
	}

	ResponseActivity struct {
		ID                         int64  `json:"id"`
		ActivityType               string `json:"activityType"`
		responseActivityAttributes `json:"activityAttribute"`
	}

	responseActivityAttributes struct {
		Name   string `json:"name"`
		Act    string `json:"act"`
		Status int    `json:"status"`
	}

	Response struct {
		Errors   `json:"errors,omitempty"`
		Data     interface{} `json:"data"`
		Messages []string    `json:"messages,omitempty"`
	}

	Errors struct {
		Code     int      `json:"code,omitempty"`
		Title    string   `json:"title,omitempty"`
		Messages []string `json:"messages,omitempty"`
	}
)
