package tariff

type Tariffs struct {
	ID           uint    `json:"id" gorm:"primary_key;column:id"`
	Name         string  `json:"name" gorm:"column:name"`
	Price        float64 `json:"price" gorm:"column:price"`
	CountDays    uint    `json:"count_days" gorm:"column:count_days"`
	Currency     string  `json:"currency" gorm:"column:currency"`
	Active       bool    `json:"active" gorm:"column:active;default:false"`
	DomainsCount uint    `json:"domains_count" gorm:"column:domains_count"`
	LinksLimit   uint    `json:"links_limit" gorm:"column:links_limit"`
	IsMy         bool    `json:"is_my" gorm:"column:is_my"`
	End          string  `json:"end"`
}

type MyTariff struct {
	ID           uint    `json:"id" gorm:"primary_key;column:id"`
	TariffId     uint    `json:"tariff_id" gorm:"column:tariff_id"`
	Name         string  `json:"name" gorm:"column:name"`
	Price        float64 `json:"price" gorm:"column:price"`
	CountDays    uint    `json:"count_days" gorm:"column:count_days"`
	Currency     string  `json:"currency" gorm:"column:currency"`
	End          string  `json:"end" gorm:"column:end"`
	DomainsCount uint    `json:"domains_count" gorm:"column:domains_count"`
	LinksLimit   uint    `json:"links_limit" gorm:"column:links_limit"`
}
