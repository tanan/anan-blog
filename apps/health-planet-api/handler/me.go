package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h Handler) Me(ctx Context) error {
	url := fmt.Sprintf("https://www.healthplanet.jp/status/innerscan.json?access_token=%s&tag=6021,6022&date=1&from=20211201000000&to=20220207000000", h.APIConfig.AccessToken)
	statusCode, resp, err := h.Client.Get(url, nil)
	if statusCode != http.StatusOK || err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	fmt.Println(string(resp))
	var scanResponse ScanResponse
	json.Unmarshal(resp, &scanResponse)
	return ctx.JSON(http.StatusOK, scanResponse)
}

type ScanResponse struct {
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
