package common

type Error struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

type Success struct {
	ID int `json:"id,omitempty"`
}

type Fair struct {
	Total int      `json:"total,omitempty"`
	Fairs []string `json:"fairs,omitempty"`
}
