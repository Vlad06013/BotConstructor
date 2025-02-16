package tariff

type Tariffs struct {
	ID          uint    `json:"id" gorm:"primary_key;column:id"`
	Name        string  `json:"name" gorm:"column:name"`
	Price       float64 `json:"price" gorm:"column:price"`
	Daylong     uint    `json:"days_long" gorm:"column:days_long"`
	Description string  `json:"description" gorm:"column:description"`
	Active      bool    `json:"active" gorm:"column:active;default:false"`
	CreatedAt   string  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   string  `json:"updated_at" gorm:"column:updated_at"`
}
