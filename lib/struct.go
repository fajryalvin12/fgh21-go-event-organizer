package lib

type Response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	PageInfo any    `json:"pageInfo,omitempty"`
	Results  any    `json:"results,omitempty"`
}

type PageInfo struct {
	TotalData int `json:"totalData"`
	TotalPage int `json:"totalPage"`
	Page      int `json:"Page"`
	Limit     int `json:"Limit"`
	Next      int `json:"Next"`
	Prev      int `json:"Prev"`
}