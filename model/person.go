package model

type Person struct {
	ID          int
	FirstName   string
	LastName    string
	Age         int
	Email       string
	Password    string  `json:"-"`
	PostAddress Address `json:"Address"`
}

type Address struct {
	HouseNo string
	City    string
}
