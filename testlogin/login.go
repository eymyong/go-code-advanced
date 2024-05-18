package testlogin

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
}

func readUsers(filename string) []user {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var u []user
	err = json.Unmarshal(b, &u)
	if err != nil {
		panic(err)
	}
	return u
}

func login(username string, password string, users []user) bool {
	for i := range users {
		v := users[i]
		if username == v.Username {
			resule := password == v.Password
			return resule
		}
	}
	return false
}

func deleteUser(users []user, userName, userPassword string) []user {
	result := []user{}
	for i := range users {
		if userName == users[i].Username {
			if userPassword == users[i].Password {
				continue
			}
		}
		result = append(result, users[i])
	}
	return result
}

func Foo2() {

	b, err := os.ReadFile("./assets/input-login3.json")
	if err != nil {
		panic(err)
	}

	var userList []user
	err = json.Unmarshal(b, &userList)
	if err != nil {
		panic(err)
	}
	fmt.Println(userList)

	newUserList := deleteUser(userList, "yong", "0000")
	fmt.Println(newUserList)

	newB, err := json.Marshal(newUserList)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./assets/output-login3.json", newB, 0664)
	if err != nil {
		panic(err)
	}

}

func Foo() {

	// users := readUsers("./assets/login.json")
	// fmt.Println(login("yong", "0022", users))

	b, err := os.ReadFile("./assets/loginmap.json")
	if err != nil {
		panic(err)
	}

	m := make(map[string]user)
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	//fmt.Println(m)

	fmt.Println(loginDataMap("yong", "admin", m))
	fmt.Println(loginDataMap("yong", "1234", m))
	fmt.Println(loginDataMap("yon", "1234", m))
	fmt.Println("======")
	fmt.Println(loginDataMap2("yong", "admin", m))
	fmt.Println(loginDataMap2("yong", "1234", m))
	fmt.Println(loginDataMap2("yon", "1234", m))

}

//==============================

func loginDataMap(username, password string, datamap map[string]user) (string, bool) {
	for i := range datamap {
		value := datamap[i]
		if username == i {
			result := password == value.Password
			return "You Can Login: ", result
		}
	}
	return "Invalid username: ", false
}

func loginDataMap2(username, password string, datamap map[string]user) (string, bool) {
	for i := range datamap {
		value := datamap[i]
		if username == i {
			if password == value.Password {
				return "You Can Login: ", true
			}
			return "Password is incorrect: ", false
		}
	}
	return "Invalid username: ", false

}
