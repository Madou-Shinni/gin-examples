package test

type CharactersFilter struct {
	Name *string `json:"name" form:"name"`
}

type SearchDTO[Filter any] struct {
	Filter *Filter `json:"filter"`
	Name   string  `form:"name"`
}
