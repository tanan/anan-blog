package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetAuth() (token string, err error) {
	redirectUri := "https://www.healthplanet.jp/success.html"
	body := []byte("client_id=" + h.APIConfig.Id + "&client_secret=" + h.APIConfig.Secret + "&redirect_uri=" + redirectUri + "&refresh_token=" + h.APIConfig.RefreshToken + "&grant_type=refresh_token")
	fmt.Println(string(body))
	url := "https://www.healthplanet.jp/oauth/token"
	header := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	statusCode, resp, err := h.Client.Post(url, header, body)
	if statusCode != http.StatusOK || err != nil {
		return "", err
	}
	fmt.Println(string(resp))
	var authResponse AuthResponse
	json.Unmarshal(resp, &authResponse)
	return authResponse.AccessToken, nil
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
