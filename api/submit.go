package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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
	formFile, err := os.Open("form.json")
	if err != nil {
		panic(err)
	}
	formBytes, err := ioutil.ReadAll(formFile)
	if err != nil {
		panic(err)
	}
	payloadMap := map[string]interface{}{}
	err = json.Unmarshal(formBytes, &payloadMap)
	if err != nil {
		panic(err)
	}
	for key, value := range payloadMap {
		payload.Set(key, fmt.Sprintf("%v", value))
	}
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
