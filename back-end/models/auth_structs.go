package models

type Http_Body struct {
	Grant_type    string `json:"grant_type"`
	Code          string `json:"code"`
	Redirect_uri  string `json:"redirect_uri"`
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
}

type Http_Body_Refresh struct {
	Grant_type    string `json:"grant_type"`
	Refresh_token string `json:"refresh_token"`
	Scope         string `json:"scope"`
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
}

type Autolab_Response struct {
	Access_token  string `json:"access_token"`
	Token_type    string `json:"token_type"`
	Expires_in    int64  `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
	Scope         string `json:"scope"`
	Created_at    int64  `json:"created_at"`
}

type Autolab_err_Response struct {
	Error             string `json:"error"`
	Error_description string `json:"error_description"`
}

type Auth_Code struct {
	Code string `json:"code"`
}
