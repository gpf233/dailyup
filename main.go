package main

import (
	"dailyup/api"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		panic(err)
	}
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	users := users{}
	json.Unmarshal(jsonBytes, &users)
	wg := sync.WaitGroup{}
	for _, user := range users {
		wg.Add(1)
		go api.Proxy(user.Username, user.Password, user.Address, &wg)
	}
	wg.Wait()
}

type users []struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
