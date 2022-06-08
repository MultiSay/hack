package model

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	ShopID     int    `json:"shop_id"`
	DriverID   int    `json:"driver_id"`
	DriverRole int    `json:"driver_role"`
	Password   string `json:"password"`
}
