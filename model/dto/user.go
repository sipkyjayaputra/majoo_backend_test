package dto

import "time"

type User struct {
	Id	int	`json:"id"`
	Name	string	`json:"name"`
	Username	string	`json:"user_name"`
	Password	string	`json:"password"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       int    `json:"created_by"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedBy       int    `json:"updated_by"`
}

type RegisterRequest struct {
	Name	string	`json:"name"`
	Username	string	`json:"user_name"`
	Password	string	`json:"password"`
}

type LoginRequest struct {
	Username	string	`json:"user_name"`
	Password	string	`json:"password"`
}

type LoginResponse struct {
	Name	string	`json:"name"`
	Username	string	`json:"user_name"`
	Token	string	`json:"token"`
}