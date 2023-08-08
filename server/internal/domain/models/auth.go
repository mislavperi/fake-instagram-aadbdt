package models

type GHCredentials struct {
	Code string `json:"code"`
}

type GHCredsReq struct {
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type GHToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GHUser struct {
	Login string `json:"login"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
