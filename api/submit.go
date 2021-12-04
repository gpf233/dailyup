package api

import (
	"io/ioutil"
	"net/http"
	"os"
)

func submit(cookies []*http.Cookie) string {
	jsonFile, err := os.Open("form.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	req, err := http.NewRequest("POST", "https://xxcapp.xidian.edu.cn/xisuncov/wap/open-report/save", jsonFile)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
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
