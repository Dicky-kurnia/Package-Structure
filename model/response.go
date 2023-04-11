package model

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"errors"`
}

type GeneralError struct {
	General string `json:"general"`
}

type NotValidImage struct {
	Image string `json:"image"`
}

type CountData struct {
	TotalData int64 `json:"total_data"`
}

type SumData struct {
	Total float64 `json:"total"`
}

type PaginationResponse struct {
	TotalPage    int32 `json:"total_page"`
	CurrentPage  int32 `json:"current_page"`
	LimitPerPage int16 `json:"limit_per_page"`
}

type GetRefResponse struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type Pagination struct {
	TotalPage   uint `json:"total_page"`
	CurrentPage uint `json:"current_page"`
}

const DateTimeLayout = "02-01-2006 15:04"

type Map map[string]interface{}
