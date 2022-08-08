package models

type User struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phone_number"`
	RegistrationDate string `json:"registration_date"`
}
