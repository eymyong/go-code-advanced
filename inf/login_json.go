package inf

import (
	"encoding/json"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginJson struct {
	data string
}

func (l loginJson) Login(username, password string) bool {
	users := []user{}
	err := json.Unmarshal([]byte(l.data), &users)
	if err != nil {
		return false
	}

	for _, user := range users {
		if user.Username == username {
			return user.Password == password
		}
	}

	return false
}

const jsonUsernamePassword = `[
	{"username": "art", "password":"1234"},
	{"username": "yong", "password":"0000"},
	{"username": "pak", "password":"admin"}
]`

const jsonEmpty = ""
