package bcgame

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/models"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/request"
	"net/http"
	"strings"
	"time"
)

var cfg Config

type Config struct {
	Host        string `json:"host"`
	Mkey        string `json:"mkey"`
	SecurityKey string `json:"securityKey"`
}

type BCPushUserBalanceChangeParams struct {
	PlayerId           string  `json:"playerId"`
	CurrencyName       string  `json:"currencyName"`
	ChangeAmount       float64 `json:"changeAmount"`
	BeforeChangeAmount float64 `json:"beforeChangeAmount"`
}

// BcGame 是对 BCGame 系统进行操作的服务
type BcGame struct {
	cfg Config
}

// NewService 创建一个新的 BcGame 实例
func NewService() *BcGame {
	var data models.TSystemDataDictionary
	mysql.DB.Debug().Model(models.TSystemDataDictionary{}).Where("name", "BCGAME").Find(&data)
	if len(data.Value) == 0 {
		logger.Error(errors.New("没有配置"))
		return nil
	}
	err := json.Unmarshal([]byte(data.Value), &cfg)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return &BcGame{cfg: cfg}
}

// 计算字符串的 MD5 哈希值
func getMD5(str string) string {
	hash := md5.Sum([]byte(str))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

// 使用 HMAC-SHA256 进行签名
func hmacSHA256Sign(message string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

// Signature 生成签名
func (s *BcGame) Signature(jsonBody, timestamp, securityKey, url string) string {
	md5Hash := getMD5(jsonBody)
	toSign := fmt.Sprintf("POST\n%s\n%s\n%s", url, timestamp, md5Hash)
	return hmacSHA256Sign(toSign, securityKey)
}

// CreateUserToken 创建用户BCGame三方登录Token
func (s *BcGame) CreateUserToken(appUserId, name, avatar string) (token string) {
	// 构建用户Token的JSON
	url := s.cfg.Host + "/merchants/webhook/player/create-code"
	tokenParams := map[string]string{
		"playerId": appUserId,
		"name":     name,
	}
	if len(avatar) > 0 {
		tokenParams["avatar"] = avatar
	}
	params, err := json.Marshal(tokenParams)
	if err != nil {
		logger.Error(err)
		return ""
	}
	Header := http.Header{}
	expTime := 5
	timestamp := fmt.Sprintf("%v", time.Now().UnixMilli()+(int64(expTime)*60*1000))
	sign := s.Signature(string(params), timestamp, s.cfg.SecurityKey, url)
	Header.Set("sign", sign)
	Header.Set("t", timestamp)
	Header.Set("mkey", s.cfg.Mkey)
	Header.Set("Content-Type", "application/json")
	post, _, err := request.Post(url, Header, string(params), nil)
	if err != nil {
		return ""
	}
	/**
	 * {
	 * "code": 0,
	 * "msg": "SUCCESS",
	 * "data": {
	 *      "token": ""
	 * }
	 */
	if len(post) > 0 {
		code := post["code"]
		if 0 == code.(float64) {
			data := post["data"].(map[string]interface{})
			token = data["code"].(string)
		} else {
			logger.Error(errors.New(post["msg"].(string)))
		}
	}
	return
}

// PushUserBalanceChange 推送用户余额变更到BCGame平台
func (s *BcGame) PushUserBalanceChange(params BCPushUserBalanceChangeParams) {
	// 将参数转换为JSON格式
	url := s.cfg.Host + "/merchants/webhook/push/balance/change"
	jsonData, err := json.Marshal(params)
	if err != nil {
		logger.Error(err)
		return
	}
	Header := http.Header{}
	expTime := 5
	timestamp := fmt.Sprintf("%v", time.Now().UnixMilli()+(int64(expTime)*60*1000))
	sign := s.Signature(string(jsonData), timestamp, s.cfg.SecurityKey, url)
	Header.Set("sign", sign)
	Header.Set("t", timestamp)
	Header.Set("mkey", s.cfg.Mkey)
	Header.Set("Content-Type", "application/json")
	post, _, err := request.Post(url, Header, string(jsonData), nil)
	if err != nil {
		logger.Error(err)
		return
	}
	if len(post) > 0 {
		code := post["code"]
		if 0 != code.(float64) {
			logger.Error(errors.New(post["msg"].(string)))
		}
	}
	logger.Info("post", post)
}

// Bg 接口定义了与BCGame系统交互的操作
type Bg interface {
	Signature(jsonBody, timestamp, securityKey, url string) string // 生成签名
	CreateUserToken(appUserId, name, avatar string)                // 创建用户BCGame三方登录Token
	PushUserBalanceChange(params BCPushUserBalanceChangeParams)    // 推送用户余额变更到BCGame平台
}
