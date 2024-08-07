package lib

type Users struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results,omitempty"`
}
