package models

type Category struct {
	Code        uint   `json:"code"`
	Description string `json:"description"`
}

func (c *Category) Create() bool {
	return false
}

func (c *Category) Read() Category {
	return Category{Code: c.Code}
}

func (c *Category) Update() Category {
	return Category{Code: c.Code}
}

func (c *Category) Delete() bool {
	return false
}
