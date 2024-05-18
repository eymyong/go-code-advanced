package inf

import "fmt"

type LoginProcessor interface {
	Login(username string, password string) bool
}

type loginDummy struct{}

func (d loginDummy) Login(username, password string) bool { return false }

func LoginTest() {
	l := loginJson{data: jsonUsernamePassword}
	err := Auth(l, "art", "1234")
	fmt.Println("login error", err)
}

func Auth(l LoginProcessor, username string, password string) error {
	if !l.Login(username, password) {
		return fmt.Errorf("login failed for username %s", username)
	}

	return nil
}

func LoginTest2() {
	l := loginViaJsonFile{filename: "./inf/login.json"}
	err := Auth(l, "art", "12345")
	fmt.Println("login error", err)
}
