package bcgame

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"goapi/pkg/logger"
	"goapi/pkg/request"
	"net/http"
	"strings"
	"time"
)

type BcGame struct {
	cfg Config
}

type Config struct {
	Host        string `json:"host"`
	Mkey        string `json:"mkey"`
	SecurityKey string `json:"securityKey"`
}

type BCPushUserBalanceChangeParams struct {
	PlayerId           string   `json:"playerId"`
	Currency           []string `json:"currency"`
	ChangeAmount       float64  `json:"changeAmount"`
	BeforeChangeAmount float64  `json:"beforeChangeAmount"`
}

// BcGameService 是对 BCGame 系统进行操作的服务
type BcGameService struct {
	cfg Config
}

// NewBcGameService 创建一个新的 BcGameService 实例
func NewBcGameService(cfg Config) *BcGameService {
	return &BcGameService{cfg: cfg}
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
func (s *BcGameService) Signature(jsonBody, timestamp, securityKey, url string) string {
	md5Hash := getMD5(jsonBody)
	toSign := fmt.Sprintf("POST\n%s\n%s\n%s", url, timestamp, md5Hash)
	return hmacSHA256Sign(toSign, securityKey)
}

// CreateUserToken 创建用户BCGame三方登录Token
func (s *BcGameService) CreateUserToken(appUserId, name, avatar string) (token string) {
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
	var Header http.Header
	expTime := 5
	timestamp := fmt.Sprintf("%v", time.Now().Unix()+(int64(expTime)*60*1000))
	sign := s.Signature(string(params), timestamp, s.cfg.SecurityKey, url)
	Header.Set("sign", sign)
	Header.Set("t", timestamp)
	Header.Set("mkey", s.cfg.Mkey)
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
		if strings.EqualFold("0", code.(string)) {
			data := post["data"].(map[string]interface{})
			token = data["code"].(string)
		} else {
			logger.Error(errors.New(post["msg"].(string)))
		}
	}
	return
}

// PushUserBalanceChange 推送用户余额变更到BCGame平台
func (s *BcGameService) PushUserBalanceChange(params BCPushUserBalanceChangeParams) {
	// 将参数转换为JSON格式
	url := s.cfg.Host + "/merchants/webhook/push/balance/change"
	jsonData, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Error marshaling balance change params:", err)
		return
	}
	var Header http.Header
	expTime := 5
	timestamp := fmt.Sprintf("%v", time.Now().Unix()+(int64(expTime)*60*1000))
	sign := s.Signature(string(jsonData), timestamp, s.cfg.SecurityKey, url)
	Header.Set("sign", sign)
	Header.Set("t", timestamp)
	Header.Set("mkey", s.cfg.Mkey)
	Header.Set("Content-Type", "application/json")
	post, _, err := request.Post(url, Header, string(jsonData), nil)
	if err != nil {
		logger.Error(err)
	}
	if len(post) > 0 {
		code := post["code"]
		if !strings.EqualFold("0", code.(string)) {
			logger.Error(errors.New(post["msg"].(string)))
		}
	}
}

// Bg 接口定义了与BCGame系统交互的操作
type Bg interface {
	Signature(jsonBody, timestamp, securityKey, url string) string // 生成签名
	CreateUserToken(appUserId, name, avatar string)                // 创建用户BCGame三方登录Token
	PushUserBalanceChange(params BCPushUserBalanceChangeParams)    // 推送用户余额变更到BCGame平台
}
