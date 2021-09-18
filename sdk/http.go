package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpPostRequest(method, reqUrl string, bodyMap map[string]interface{}) ([]byte, error) {
	res, err := json.Marshal(bodyMap)
	if err != nil {
		fmt.Println(res)

	}
	req, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(res))
	if err != nil {

	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	reqBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return reqBody,err
}

