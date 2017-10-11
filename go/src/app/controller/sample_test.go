package controller

import (
	"app/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenazn/goji/web"
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

func TestSample(t *testing.T) {
	is := assert.New(t)
	m := web.New()
	SetUpSample(m)
	ts := httptest.NewServer(m)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/api/v1/sample")
	is.NoError(err)
	v := model.Sample{}
	status := JsonParseResponse(res, &v)
	is.Equal(http.StatusOK, status)
	is.Equal("title", v.Title)
	is.Equal("name", v.Name)
}
