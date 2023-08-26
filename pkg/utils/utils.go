package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		// fmt.Println(body)
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
