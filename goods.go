package config

// Goods configuration structure
type Goods struct {
	// Item - good code for sales product
	Item int

	// Shipping - good code for shipping product
	Shipping *int

	// Tax - good code for tax product
	Tax *int

	// Gift - good code for gift wrap product
	Gift *int
}
