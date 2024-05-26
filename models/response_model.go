package models

type Repsonse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func RepsonseModel() Repsonse {
	return Repsonse{}
}
