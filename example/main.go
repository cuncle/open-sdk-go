package main

import (
	"fmt"
	"net/http"
	"open-sdk-go/client"
	"strconv"
)

func getAccessToken(authorize_type, client_id, client_secret, grant_id string, refresh bool) {
	url := "https://open.youzanyun.com/auth/token"
	res := map[string]string{
		"authorize_type": authorize_type,
		"client_id":      client_id,
		"client_secret":  client_secret,
		"grant_id":       grant_id,
		"refresh":        strconv.FormatBool(refresh),
	}
	result := client.HttpRequest(http.MethodPost, url, res)
	fmt.Println(result)
}

func main() {
	getAccessToken("silent", "xxxx", "xxxx", "kdt_id", false)
}