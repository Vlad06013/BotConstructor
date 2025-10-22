package tariff

type Tariffs struct {
	ID                  uint    `json:"id" gorm:"primary_key;column:id"`
	Name                string  `json:"name" gorm:"column:name"`
	Price               float64 `json:"price" gorm:"column:price"`
	CountDays           uint    `json:"count_days" gorm:"column:count_days"`
	Currency            string  `json:"currency" gorm:"column:currency"`
	Active              bool    `json:"active" gorm:"column:active;default:false"`
	DomainsCount        uint    `json:"domains_count" gorm:"column:domains_count"`
	DomainsUnlimited    bool    `json:"domains_unlimited" gorm:"column:domains_unlimited"`
	LinksForDomainCount uint    `json:"links_for_domain_count" gorm:"column:links_for_domain_count"`
}

type MyTariff struct {
	ID                  uint    `json:"id" gorm:"primary_key;column:id"`
	TariffId            uint    `json:"tariff_id" gorm:"column:tariff_id"`
	Name                string  `json:"name" gorm:"column:name"`
	Price               float64 `json:"price" gorm:"column:price"`
	CountDays           uint    `json:"count_days" gorm:"column:count_days"`
	Currency            string  `json:"currency" gorm:"column:currency"`
	End                 string  `json:"end" gorm:"column:end"`
	DomainsCount        uint    `json:"domains_count" gorm:"column:domains_count"`
	DomainsUnlimited    bool    `json:"domains_unlimited" gorm:"column:domains_unlimited"`
	LinksForDomainCount uint    `json:"links_for_domain_count" gorm:"column:links_for_domain_count"`
}
