package goUtils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func reqJsonToObject(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else {
		json.Unmarshal(content, &v)
		return nil
	}
}
