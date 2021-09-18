package sdk

import (
"encoding/json"
"net/http"
)

const UrlAuthToken string = "https://open.youzanyun.com/auth/token"

type RequestTokenPayload struct {
	AuthorizeType string `json:"authorize_type"`
	ClientID      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	GrantID string `json:"grant_id"`
	Refresh bool `json:"refresh"`

}

func (req *RequestTokenPayload) toMap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["authorize_type"]=req.AuthorizeType
	m["client_id"]=req.ClientID
	m["client_secret"]=req.ClientSecret
	m["grant_id"] = req.GrantID
	m["refresh"] =  req.Refresh
	return
}


// 获取自用型AccessToken响应参数结构体
type GetSelfTokenResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data struct {
		AccessToken string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		Expires     int64  `json:"expires"`
		Scope       string `json:"scope"`
	} `json:"data"`
}

func GetAccessToken(request RequestTokenPayload)(response GetSelfTokenResponse, err error) {
	body,err:= HttpPostRequest(http.MethodPost, UrlAuthToken, request.toMap())
	err=json.Unmarshal(body, &response)
	return
}


