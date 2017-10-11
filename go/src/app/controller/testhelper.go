package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(respBody), res.StatusCode
}

func JsonParseResponse(res *http.Response, dst interface{}) int {
	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(respBody, dst)
	if err != nil {
		panic(err)
	}
	return res.StatusCode
}
