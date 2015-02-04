package goUtils

import (
	"bytes"
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

func PostJson(url string, params interface{}) []byte {
	jsonStr, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	return content
}
