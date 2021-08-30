package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(method, reqUrl string, bodyMap map[string]string) (result string) {
	res, err := json.Marshal(bodyMap)
	if err != nil {
		fmt.Println(res)
		return
	}
	req, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(res))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	reqBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(reqBody)
}
