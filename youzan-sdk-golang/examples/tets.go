package main

import (
	"fmt"
	"open-sdk-go/sdk"
	yzclnt "youzan-sdk-golang/sdk"
)
func getAccessTokenSelf(clientID, clientSecret, grantID string)  (resp sdk.GetSelfTokenResponse, err error){
	return sdk.GetAccessToken(sdk.RequestTokenPayload{
		AuthorizeType: "silent",
		ClientID:clientID,
		ClientSecret:clientSecret,
		GrantID:grantID,
		Refresh:false})
}
func main() {
resp,_:=getAccessTokenSelf("your_client_id","your_client_secret","your_kdt_id")
clnt:=yzclnt.New(resp.Data.AccessToken)
params := make(map[string]interface{})
params["mobile"] = "13998981212"
params["open_user_id"]="13998981212"
params["country_code"]="+86"
r,_:=clnt.ApiPost("youzan.user.platform.import","1.0.0",params)
fmt.Println(string(r))
}

