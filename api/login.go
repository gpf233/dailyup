package api

import (
	"net/http"
	"net/url"
)

func login(username string, password string) []*http.Cookie {
	resp, err := http.PostForm("https://xxcapp.xidian.edu.cn/uc/wap/login/check", url.Values{"username": {username}, "password": {password}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.Cookies()
}
