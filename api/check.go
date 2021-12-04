package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func check(cookies []*http.Cookie) (string, error) {
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
	profile := profile{}
	json.Unmarshal(bodyBytes, &profile)
	if !strings.Contains(string(bodyBytes), "操作成功") {
		return profile.D.Userinfo.Realname, errors.New("login error")
	}
	return profile.D.Userinfo.Realname, nil
}
