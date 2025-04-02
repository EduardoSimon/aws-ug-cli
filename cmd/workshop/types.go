package workshop

// Product represents a product in the catalog
type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Category    string   `json:"category"`
	Brand       string   `json:"brand"`
	Stock       int      `json:"stock"`
	Rating      float64  `json:"rating"`
	Tags        []string `json:"tags"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
