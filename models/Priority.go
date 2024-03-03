package models

type Priority struct {
	Code        uint   `json:"code"`
	Description string `json:"description"`
}

func (p *Priority) Create() bool {
	return false
}

func (p *Priority) Read() Priority {
	return Priority{Code: p.Code}
}

func (p *Priority) Update() Priority {
	return Priority{Code: p.Code}
}

func (p *Priority) Delete() bool {
	return false
}
