package models

type User struct {
	Uid          uint64 `json:"uid" gorm:"primaryKey"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Contact_info uint64 `json:"contact_info"`
}
