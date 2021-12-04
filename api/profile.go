package api

type profile struct {
	E int    `json:"e"`
	M string `json:"m"`
	D struct {
		Date string `json:"date"`
		Info struct {
			Ymtys      string `json:"ymtys"`
			Sfzx       string `json:"sfzx"`
			Tw         string `json:"tw"`
			Area       string `json:"area"`
			City       string `json:"city"`
			Province   string `json:"province"`
			Address    string `json:"address"`
			GeoAPIInfo string `json:"geo_api_info"`
			Sfcyglq    string `json:"sfcyglq"`
			Sfyzz      string `json:"sfyzz"`
			Date       string `json:"date"`
			Flag       int    `json:"flag"`
			UID        string `json:"uid"`
			Creator    string `json:"creator"`
			Created    int    `json:"created"`
			ID         int    `json:"id"`
		} `json:"info"`
		Ontime   bool `json:"ontime"`
		Realonly bool `json:"realonly"`
		Userinfo struct {
			UID          string `json:"uid"`
			Realname     string `json:"realname"`
			Sex          string `json:"sex"`
			Email        string `json:"email"`
			Mobile       string `json:"mobile"`
			Weixin       string `json:"weixin"`
			Qq           string `json:"qq"`
			Csrq         string `json:"csrq"`
			Jg           string `json:"jg"`
			Mz           string `json:"mz"`
			Zjhlx        string `json:"zjhlx"`
			Zjh          string `json:"zjh"`
			Hyzk         string `json:"hyzk"`
			Jkzk         string `json:"jkzk"`
			Gj           string `json:"gj"`
			Xx           string `json:"xx"`
			Whcd         string `json:"whcd"`
			Zgxl         string `json:"zgxl"`
			Xyzj         string `json:"xyzj"`
			Zzmm         string `json:"zzmm"`
			Photo        string `json:"photo"`
			Avatar       string `json:"avatar"`
			ActiveRoleID string `json:"active_role_id"`
			IsDel        string `json:"is_del"`
			Role         struct {
				Roleid     int    `json:"roleid"`
				Number     string `json:"number"`
				Identity   string `json:"identity"`
				IdentityID int    `json:"identity_id"`
				Sfzx       int    `json:"sfzx"`
			} `json:"role"`
			Departs struct {
				Num410413891 string `json:"410413891"`
			} `json:"departs"`
			Age  int    `json:"age"`
			Type string `json:"type"`
		} `json:"userinfo"`
		Setting struct {
			Title     string `json:"title"`
			Desc      string `json:"desc"`
			Area      string `json:"area"`
			Sendmsg   string `json:"sendmsg"`
			Img       string `json:"img"`
			Copyright string `json:"copyright"`
			Notifywid string `json:"notifywid"`
			Sms       string `json:"sms"`
			City      string `json:"city"`
			NoticeNum string `json:"notice_num"`
			ModifyHis string `json:"modify_his"`
		} `json:"setting"`
	} `json:"d"`
}
