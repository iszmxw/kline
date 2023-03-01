package goadjust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/pkg/config"
	"goapi/pkg/mysql"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	// ErrDeviceNotFound is returned when adjust did not recognise the device
	ErrDeviceNotFound = errors.New("device not found")
)

// Response represents the response from an track request.
type Response struct {
	Status       string `json:"status"`
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

func Get(url, authorization string) (resp *Response, err error) {
	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	// 增加header选项
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authorization))
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, response.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
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

func submitEvent(eventParams EventParams, nation, partnerCode string) {
	params := cmap.New().Items()
	if checkSendLog(params) {
		return
	}
	accountCode := ""
	if len(eventParams.EventBindId) > 0 && strings.Contains(strings.ToLower(eventParams.DeviceBrand), "apple") {
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

	req := make(url.Values)
	// Add client fields to request
	req.Add("s2s", "1")
	req.Add("app_token", "majbv8e0m8e8")
	req.Add("event_token", "y0ec7c")
	req.Add("gps_adid", "482bd535-7fcd-49ea-9678-9b9d15e81737")
	req.Add("callback_params", "{\"userId\":\"7FF235ACCD4\"}")
	//req.Add("idfa", "D2CADB5F-410F-4963-AC0C-2A78534BDF1E")
	u := "https://s2s.adjust.com/event?" + req.Encode()
	send, err := Get(u, "2dc8c77d0503c9aab03cdeb6ebc5849b")
	if err != nil {
		return
	}
	fmt.Println(send)
}

// 查询用户发送成功记录，有记录返回true
func checkSendLog(params map[string]interface{}) bool {
	var count int64
	where := cmap.New().Items()
	where["event_name"] = ""
	where["app_user_id"] = ""
	where["email"] = ""
	where["status"] = ""
	mysql.DB.Debug().Model(models.AppUserAdjustSendLog{}).Where(where).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
