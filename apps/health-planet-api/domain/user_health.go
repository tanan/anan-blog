package domain

type UserHealth struct {
	BirthDate string       `json:"birth_date"`
	Height    string       `json:"height"`
	Sex       string       `json:"sex"`
	Data      []HealthData `json:"data"`
}

type HealthData struct {
	date    string `json:"date"`
	keydata string `json:"keydata"`
	model   string `json:"model"`
	tag     string `json:"tag"`
}
