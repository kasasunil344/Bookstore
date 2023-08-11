package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return 
	}
	return
}