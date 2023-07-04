package goadjust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	// ErrDeviceNotFound is returned when adjust did not recognise the device
	ErrDeviceNotFound = errors.New("device not found")
)

type EventTokenStrObj struct {
	Data struct {
		Register         string `json:"register"`
		FirstPurchase    string `json:"firstPurchase"`
		Purchase         string `json:"purchase"`
		Trade            string `json:"trade"`
		Withdraw         string `json:"withdraw"`
		CompleteTutorial string `json:"completeTutorial"`
		SkipTutorial     string `json:"skipTutorial"`
		FirstWithdraw    string `json:"firstWithdraw"`
	} `json:"eventTokenStrObj"`
}

// Response represents the response from a track request.
type Response struct {
	Status       string `json:"status"`
	Error        string `json:"error"`
	TrackerToken string `json:"tracker_token"`
	TrackerName  string `json:"tracker_name"`
	Network      string `json:"network"`
	Country      string `json:"country"`
}

type EventParams struct {
	EventName        string `json:"eventName"`        // 事件名称 (必传参数)
	AppUserId        string `json:"appUserId"`        // 用户id (必传参数)
	IsRepeatedSubmit string `json:"isRepeatedSubmit"` // 是否可重复提交 1是0否 (选传，不传时默认不可重复提交)
	Mobile           string `json:"mobile"`           // 客户手机号
	EventBindId      string `json:"eventBindId"`      // 第三方绑定的id，例如：adjust:GpsAdid
	Email            string `json:"email"`            // 邮箱
	FlowJsonString   string `json:"flowJsonString"`   // 第三方参数json
	DeviceBrand      string `json:"deviceBrand"`      // 用户设备品牌
}

type AdjustParams struct {
	AppToken      string `json:"appToken"`
	Authorization string `json:"authorization"`
	Url           string `json:"url"`
}

type FlowJson struct {
	Adid string `json:"adid"`
	Idfa string `json:"idfa"`
	Idfv string `json:"idfv"`
}

func Get(url, authorization string) (resp *Response, err error) {
	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("POST", url, nil)
	// 增加header选项
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authorization))
	if err != nil {
		logger.Error(err)
		return
	}
	//处理返回结果
	response, _ := client.Do(request)
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, response.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}(response.Body)
	switch {
	case strings.Contains(buf.String(), "Device not found"):
		return nil, ErrDeviceNotFound
	case strings.HasPrefix(buf.String(), "Event failed"):
		msg := strings.TrimSpace(buf.String())
		return nil, errors.New(msg)
	}
	// Unmarshal response
	if err := json.NewDecoder(buf).Decode(&resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// 提交 adjust 事件

func SubmitEvent(eventParams EventParams, nation, partnerCode string) {
	params := cmap.New().Items()
	params["event_name"] = eventParams.EventName
	params["app_user_id"] = eventParams.AppUserId
	params["email"] = eventParams.Email
	if checkSendLog(params) {
		return
	}
	accountCode := ""
	if len(eventParams.DeviceBrand) > 0 && strings.Contains(strings.ToLower(eventParams.DeviceBrand), "apple") {
		accountCode = "apple"
	} else {
		accountCode = "android"
	}
	var collaboratorParams models.CollaboratorParams
	mysql.DB.Debug().Model(models.CollaboratorParams{}).Where(map[string]interface{}{
		"item_code":    config.GetString("app.item_code"),
		"partner_code": partnerCode,
		"account_code": accountCode,
		"nation":       nation,
	}).Find(&collaboratorParams)
	if collaboratorParams.Id <= 0 {
		logger.Info("参数未配置")
		return
	}
	var systemApiRule models.SystemApiRule
	mysql.DB.Debug().Model(models.SystemApiRule{}).Where(map[string]interface{}{
		"is_start":  "1",
		"rule_code": "SEND_ADJUST_HIDDEN_RULE",
		"item_code": config.GetString("app.item_code"),
		"api_code":  accountCode,
		"nation":    nation,
	}).Find(&systemApiRule)
	if systemApiRule.Id <= 0 {
		logger.Info("参数未配置")
		return
	}
	/**解析adjust配置参数*/
	var adjustParams AdjustParams
	err := json.Unmarshal([]byte(collaboratorParams.ParamJson), &adjustParams)
	if err != nil {
		logger.Error(err)
		return
	}
	var ruleJson EventTokenStrObj
	err = json.Unmarshal([]byte(systemApiRule.RuleJson), &ruleJson)
	if err != nil {
		logger.Error(err)
		return
	}
	eventToken := ruleJson.Data.Trade
	req := make(url.Values)
	req.Add("s2s", "1")
	req.Add("app_token", adjustParams.AppToken)
	req.Add("event_token", eventToken)
	adid := ""
	if len(eventParams.EventBindId) == 0 && len(eventParams.FlowJsonString) > 0 {
		var flowJson FlowJson
		err = json.Unmarshal([]byte(eventParams.FlowJsonString), &flowJson)
		if err == nil {
			if len(flowJson.Adid) > 0 {
				adid = flowJson.Adid
				req.Add("adid", flowJson.Adid)
			}
			if len(flowJson.Idfa) > 0 {
				req.Add("idfa", flowJson.Idfa)
			}
			if len(flowJson.Idfv) > 0 {
				req.Add("idfv", flowJson.Idfv)
			}
		} else {
			logger.Error(err)
		}
	} else {
		req.Add("gps_adid", eventParams.EventBindId)
	}
	if len(eventParams.AppUserId) > 0 {
		req.Add("callback_params", fmt.Sprintf(`{"userId":"%s"}`, eventParams.AppUserId))
	}
	u := adjustParams.Url + "?" + req.Encode()
	resp, err := Get(u, adjustParams.Authorization)
	if err != nil {
		logger.Error(err)
		return
	}
	// 保存请求 adjust 结果信息
	createData := models.AppUserAdjustSendLog{
		ItemCode:   config.GetString("app.item_code"),
		AppUserId:  eventParams.AppUserId,
		EventName:  eventParams.EventName,
		CreateDate: time.Now(),
		EventToken: eventToken,
		AppToken:   adjustParams.AppToken,
		GpsAdid:    eventParams.EventBindId,
		Mobile:     eventParams.Mobile,
		UpdateDate: time.Now(),
		Email:      eventParams.Email,
		Adid:       adid,
		Status:     0,
	}
	if resp != nil {
		if len(resp.Status) > 0 {
			createData.Status = 1
			createData.ResultMessage = resp.Status
		} else if len(resp.Error) > 0 {
			createData.Status = 0
			createData.ResultMessage = resp.Error
		}
	}
	err = mysql.DB.Debug().Model(models.AppUserAdjustSendLog{}).Create(&createData).Error
	if err != nil {
		logger.Error(err)
	}
}

// 查询用户发送成功记录，有记录返回true
func checkSendLog(params map[string]interface{}) bool {
	var count int64
	where := cmap.New().Items()
	where["event_name"] = params["event_name"]
	if len(params["app_user_id"].(string)) > 0 {
		where["app_user_id"] = params["app_user_id"]
	}
	if len(params["email"].(string)) > 0 {
		where["email"] = params["email"]
	}
	where["status"] = 1
	mysql.DB.Debug().Model(models.AppUserAdjustSendLog{}).Where(where).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
