package model

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserID     int    `json:"user_id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	ShopID     int    `json:"shop_id"`
	DriverID   int    `json:"driver_id"`
	DriverRole int    `json:"driver_role"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

type TokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token,omitempty"`  // AccessToken
	RefreshToken string `json:"refresh_token,omitempty"` // RefreshToken
}
