package inf

import (
	"encoding/json"
	"os"
)

type loginViaJsonFile struct {
	filename string // read from this filename
}

func (l loginViaJsonFile) Login(username, password string) bool {
	b, err := os.ReadFile(l.filename)
	if err != nil {
		return false
	}
	var data map[string]string
	err = json.Unmarshal(b, &data)
	if err != nil {
		return false
	}

	for k, v := range data {
		if username == k {
			if password == v {
				return true
			}
		}
	}
	return false
}
