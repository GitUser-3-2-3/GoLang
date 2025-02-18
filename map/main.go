package main

import (
	"errors"
	"fmt"
)

type user struct {
	username    string
	phoneNumber int
}

func getUserMap(users []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(users) != len(phoneNumbers) {
		return nil, errors.New("length of users and phoneNumbers slices should be equal")
	}
	for i := 0; i < len(users); i++ {
		username := users[i]
		phoneNumber := phoneNumbers[i]
		userMap[username] = user{username: username, phoneNumber: phoneNumber}
	}
	return userMap, nil
}

func main() {
	users := []string{"user1", "user2", "user3"}
	phoneNumbers := []int{1234567890, 2345678901, 3456789012}
	userMap, err := getUserMap(users, phoneNumbers)
	fmt.Println(userMap)
	if err != nil {
		panic(err)
	}
	for key, value := range userMap {
		fmt.Println("Username: ", key, " Phone Number: ", value.phoneNumber)
	}
}
