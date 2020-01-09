package config

import (
	"time"
)

// Butent - configuration for software Butent
type Butent struct {
	// Letter for table [apyvarta], field [raide]
	Letter string

	// HomeCountry checking for home operaion
	HomeCountry string

	// Operations is 3 values:
	// home   - operation for country set in HomeCountry
	// euro   - operations for Euro zone, cheking countries in table [salys]
	//          and have mode [e]
	// other  - operation for other countries, do not have mode [e]
	Operations map[string]string

	// Warehause - Value for table [apyvarta], field [tiek_sand]
	Warehouse string

	// Isaf - Integer value for table [apyvarta], field [isaf], zero or nil for disabled
	Isaf *int

	// Date - configure date using in projects
	Date *string

	// Client - Value for table [apyvarta], field [gavejas].
	Client *string

	// Good - Values for table [apyv_gr], field [preke]
	Goods

	// UserID - Value for tables field [inp_user]
	UserID int

	// Vat - Value for table [apyv_gr], field [pvm_stat]
	Vat Vat
}

// GetGoodItem return Item code from configuration
func (b *Butent) GetGoodItem() int {
	return b.Goods.Item
}

// GetShippItem return Shipping code from configuration
func (b *Butent) GetShippItem() *int {
	return b.Goods.Shipping
}

// GetTaxItem return Tax code from configuration
func (b *Butent) GetTaxItem() *int {
	return b.Goods.Tax
}

// GetGiftItem return Gift code from configuration
func (b *Butent) GetGiftItem() *int {
	return b.Goods.Gift
}

// GetDate return date by Date configuration, of not set in configuration return t time
// Otherwise use startday, endday or by default now (time.Now())
func (b *Butent) GetDate(t time.Time) time.Time {
	if b.Date == nil {
		return t
	}

	t = time.Now()
	year, month, day := t.Date()

	switch *b.Date {
	case "startday":
		t = time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	case "endday":
		t = time.Date(year, month, day, 23, 59, 59, 0, t.Location())
	}

	return t
}
