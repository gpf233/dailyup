package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func submit(cookies []*http.Cookie, address string) string {
	payload := url.Values{}
	if address == "s" {
		payload.Set("area", "陕西省 西安市 长安区")
	} else if address == "n" {
		payload.Set("area", "陕西省 西安市 雁塔区")
	} else {
		panic(errors.New("address error"))
	}
	payload.Set("tw", "2")
	payload.Set("ymtys", "0")
	payload.Set("sfzx", "1")
	payload.Set("sfcyglq", "0")
	payload.Set("sfyzz", "0")
	payload.Set("qtqk", "无")
	req, err := http.NewRequest("POST", "https://xxcapp.xidian.edu.cn/xisuncov/wap/open-report/save", strings.NewReader(payload.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
	return string(bodyBytes)
}
