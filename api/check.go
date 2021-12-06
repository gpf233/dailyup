package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func check(cookies []*http.Cookie, username string) string {
	req, err := http.NewRequest("GET", "https://xxcapp.xidian.edu.cn/xisuncov/wap/open-report/index", nil)
	if err != nil {
		panic(err)
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if !strings.Contains(string(bodyBytes), "操作成功") {
		log.Panicln(errors.New(username +  " login error"))
	}
	profile := profile{}
	err = json.Unmarshal(bodyBytes, &profile)
	if err != nil {
		panic(err)
	}
	return profile.D.Userinfo.Realname
}
