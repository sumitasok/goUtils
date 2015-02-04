package goUtils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReqJsonToObject(req *http.Request, v interface{}) error {
	defer req.Body.Close()

	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	} else {
		json.Unmarshal(content, &v)
		return nil
	}
}

func RespJsonToObject(req *http.Response, v interface{}) error {
	defer req.Body.Close()

	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	} else {
		json.Unmarshal(content, &v)
		return nil
	}
}
