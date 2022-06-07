package api

import (
	"encoding/json"
	"errors"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	api2captcha "goapi/pkg/recaptcha"
	"goapi/pkg/redis"
	"goapi/pkg/request"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-olymptrade/olymptrade/config"
	"net/http"
	"strings"
	"sync"
)

// 登录结构

type Login struct {
	List    []string `json:"list"`
	Captcha struct {
		Type string `json:"type"`
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	} `json:"captcha"`
	Data struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		FpId       string `json:"fp_id"`
		HardwareId string `json:"hardware_id"`
	} `json:"data"`
}

type SocketSession struct {
	Login   Login     // 登录结构
	Once    sync.Once // 锁
	IsLogin bool      // 是否已经登录
}

// 获取登录token

func (ss *SocketSession) GetToken(refresh bool) (error, string) {
	// 读取 session，如果未配置session进行邮件提醒
	redisLoginKey := "kline:token:olymptrade:login"
	// 刷新token
	token, _ := redis.Client.Get(config.RedisKey)
	// token失效，或者强制刷新token,就重新登录获取token
	if len(token) <= 0 || true == refresh {
		if ss.IsLogin {
			return errors.New("已登录，请不要重复操作"), ""
		}
		isLogin, _ := redis.Client.Get(redisLoginKey)
		// 一分钟之内登录过
		if len(isLogin) > 0 {
			return errors.New("一分钟之内已经进行登录操作，请稍后再试"), ""
		}
		// 登录
		ss.Once.Do(func() {
			ss.IsLogin = true
			_, err := redis.Client.Add(redisLoginKey, "true", 60) // 60秒之内只能登录一次
			if err != nil {
				logger.Error(err)
			}
			token = login()
			_, err = redis.Client.Add(config.RedisKey, token, 60*60*24*10) // 设置十天过期
			if err != nil {
				logger.Error(err)
			}
		})
	}
	// 返回 token
	return nil, token
}

func login() string {
	var lData Login
	var token string
	// list
	lData.List = []string{"data"}
	// captcha
	lData.Captcha.Type = "recaptcha"
	lData.Captcha.Data.Token = recaptcha()
	// data
	lData.Data.HardwareId = "hi_AXZujaPKIo1wlwthz1Ab3B2Oxa3CEtp8"
	if conf.GetString("app.env") == "company_pro" {
		lData.Data.Email = "open@54zm.com"
		lData.Data.Password = "wap.54zm.com"
		lData.Data.FpId = "32c2d37cf06c08b2b42c40331bc906af"
	} else {
		lData.Data.Email = "435311939@qq.com"
		lData.Data.Password = "18219116520tdw"
		lData.Data.FpId = "73ccb7266cf5f71cae07b114c087814e"
	}
	str, err := json.Marshal(lData)
	if err != nil {
		logger.Error(err)
		return ""
	}
	resp, header, err := request.Post("https://api.olymptrade.com/v1/user/login-by-password", http.Header{
		"X-Request-Project": {"otp"},
		"content-type":      {"application/json"},
		"origin":            {"https://olymptrade.com"},
		"referer":           {"https://olymptrade.com/"},
		"User-Agent":        {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"},
	}, string(str), kline_common.GetProxy(config.KlineSource))
	if err != nil {
		logger.Error(err)
	}
	if resp == nil {
		logger.Info(resp)
	}
	token = exportSession(header)
	logger.Info(token)
	return token
}

// 获取 recaptcha
func recaptcha() string {
	client := api2captcha.NewClient("db18edb60341ce0c0d93d8905ed3cdaf")
	captcha := api2captcha.ReCaptcha{
		SiteKey:   "6Lc5ACgTAAAAABL3s8j9VkkUn4Engv4QtdCLd9qI",
		Url:       "https://olymptrade.com/home",
		Invisible: true,
		Action:    "verify",
	}
	req := captcha.ToRequest()
	code, err := client.Solve(req)
	if err != nil {
		logger.Error(err)
	}
	return code
}

// 导出

func exportSession(header http.Header) string {
	var result string
	for _, item := range header.Values("Set-Cookie") {
		if strings.Contains(item, "session") {
			result = item
		}
	}
	if len(result) > 0 {
		arr := strings.Split(result, ";")
		for _, item := range arr {
			if strings.Contains(item, "session") {
				result = item
			}
		}
	}
	return result
}
