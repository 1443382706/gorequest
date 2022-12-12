package request

import (
	"net/http"
)

func GetMore(filter map[string]string, request *http.Request) map[string]string{
	params := make(map[string]string)
	request.ParseForm()
	for k, v := range filter {
		value, flag := request.Form[k]
		if flag {
			params[k] = value[0]
		}else{
			params[k] = v
		}
	}

	return params
}