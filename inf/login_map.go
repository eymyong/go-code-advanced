package inf

var mapUsernamePassword = map[string]string{
	"art":  "1234",
	"yong": "0000",
	"pak":  "admin",
}

type loginMap struct{}

func (l loginMap) Login(username, password string) bool {
	pass, ok := mapUsernamePassword[username]
	if !ok {
		return false
	}

	return password == pass
}
