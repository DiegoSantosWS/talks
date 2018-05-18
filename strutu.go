//Products Struct of products
type Products struct {
	ID          int     `json:"id" db:"id"`
	Nome        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Stock       float64 `json:"stock" db:"stock"`
	Width       float64 `json:"width" db:"width"`
	Height      float64 `json:"heigth" db:"height"`
	Amount      float64 `json:"amount" db:"amount"`
	Weight      float64 `json:"weight" db:"weight"`
	Price       float64 `json:"price" db:"price"`
	Discount    float64 `json:"discount" db:"discount"`
	Promotion   float64 `json:"promotion" db:"promotion"`
}
