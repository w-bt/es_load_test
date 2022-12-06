package models

type Item struct {
	Name          string   `json:"name,omitempty"`
	BrandName     string   `json:"brand_name,omitempty"`
	CategoryNames []string `json:"category_names,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Price         int64    `json:"price,omitempty"`
	Code          string   `json:"code,omitempty"`
	Score         *float64 `json:"-"`
}
