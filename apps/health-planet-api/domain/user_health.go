package domain

type UserHealth struct {
	BirthDate string       `json:"birth_date"`
	Height    string       `json:"height"`
	Sex       string       `json:"sex"`
	Data      []HealthData `json:"data"`
}

type HealthData struct {
	Date    string `json:"date"`
	Keydata string `json:"keydata"`
	Model   string `json:"model"`
	Tag     string `json:"tag"`
}
