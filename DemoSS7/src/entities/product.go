package entities

type Product struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Category Category  `json:"category"`
	Comments []Comment `json:"comments"`
}
