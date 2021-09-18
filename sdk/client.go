package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type YouzanyunClient interface {
	Conf() *YouzanyunConf
	WithConf(conf *YouzanyunConf) YouzanyunClient
	WithHttp(http *http.Client) YouzanyunClient
	HttpClnt() *http.Client
	ApiPost(ApiName,ApiVersion string,body map[string]interface{})([]byte,error)
	Close()
}
// youzanyunconf 基础

type YouzanyunConf struct {
	AuthHost,ApiHost,AuthrizeType,ClientID,ClientSecret,GrantID string
	Refresh bool
	Http *HttpConf
}
type HttpConf struct {
	Timeout   time.Duration
	KeepAlive time.Duration
	MaxIdleConns        int
	IdleConnTimeout     time.Duration // second
	TLSHandshakeTimeout time.Duration
}


type HttpClnt struct {
	accesstoken string
	conf   YouzanyunConf
	http   *http.Client
}

func (clnt *HttpClnt) Conf() *YouzanyunConf {
	return &clnt.conf
}
func (clnt *HttpClnt) HttpClnt() *http.Client {
	return clnt.http
}
func(clnt *HttpClnt) Accesstoken() string{
	return clnt.accesstoken
}

var DefOnlineConf = &YouzanyunConf{
	ApiHost: "https://open.youzanyun.com/api/",
	Http: &HttpConf{Timeout: 30, KeepAlive: 30, MaxIdleConns: 100, IdleConnTimeout: 30, TLSHandshakeTimeout: 10},
}

func New(accesstoken string) YouzanyunClient {
	return &HttpClnt{
		accesstoken,
		*DefOnlineConf,
		createHttp(DefOnlineConf),
	}
}

func createHttp(conf *YouzanyunConf) *http.Client {
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   conf.Http.Timeout * time.Second,
			KeepAlive: conf.Http.KeepAlive * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:        conf.Http.MaxIdleConns,
		IdleConnTimeout:     conf.Http.IdleConnTimeout * time.Second,
		TLSHandshakeTimeout: conf.Http.TLSHandshakeTimeout * time.Second,
	}
	return &http.Client{Transport: tr}
}
func (clnt *HttpClnt) WithConf(conf *YouzanyunConf) YouzanyunClient {
	if conf != nil {
		clnt.conf = *conf
	}
	clnt.http = createHttp(&clnt.conf)
	return clnt
}

func (clnt *HttpClnt) WithHttp(http *http.Client) YouzanyunClient {
	if http != nil {
		clnt.http = http
	}
	return clnt
}


func (clnt *HttpClnt) ApiPost(ApiName,ApiVersion string,body map[string]interface{}) ([]byte, error) {
	Requesturl:=clnt.conf.ApiHost + ApiName + "/" + ApiVersion +"?access_token=" + clnt.accesstoken
	res, err := json.Marshal(body)
	if err != nil {
		fmt.Println(res)

	}
	rep, err := http.NewRequest("POST", Requesturl, bytes.NewBuffer(res))
	if err != nil {
		return nil, err
	}
	rep.Header.Add("Content-Type", "application/json")
	resp,err:=clnt.http.Do(rep)
	if err != nil {
		panic(err)
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	return reqBody,err
}



func (clnt *HttpClnt) Close() {
	//
}



