package models

type User struct {
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
	Address          string `json:"address" bson:"address"`
	PhoneNumber      string `json:"phone_number" bson:"phone_number"`
	RegistrationDate string `json:"registration_date" bson:"registration_date"`
}
