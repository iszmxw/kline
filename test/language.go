package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/robertkrimen/otto"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"io/ioutil"
	"os"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()

	vm := otto.New()
	file, err := os.Open("D:\\www\\kline\\test\\language.js")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	jsByte, err := ioutil.ReadAll(file)
	vm.Run(string(jsByte))
	commonFunc(vm, "common")
	newsPageFunc(vm, "newsPage")
	newsPageInfoFunc(vm, "newsPageInfo")
	bonusActivitiesFunc(vm, "bonusActivities")
	paymentInfoFunc(vm, "inviteFriends")
	paymentInfoFunc(vm, "tradingGuide")
	paymentInfoFunc(vm, "lossReturn")
	paymentInfoFunc(vm, "paymentInfo")
	paymentInfoFunc(vm, "rechargeInfo")
	paymentInfoFunc(vm, "troubleshoot")
	termsOfServiceFunc(vm, "termsOfService")
	paymentInfoFunc(vm, "userAgreement")
	paymentInfoFunc(vm, "tradePage")
}

func commonFunc(vm *otto.Otto, name string) {
	obj, _ := vm.Get(name)
	common := obj.Object()
	languageCodes := common.Keys()
	for _, languageCode := range languageCodes {
		data, _ := common.Get(languageCode)
		keys := data.Object().Keys()
		for _, key := range keys {
			val, _ := data.Object().Get(key)
			newKey := fmt.Sprintf("%s.%s", name, key)
			insertData(languageCode, newKey, val.String())
		}
	}
}

func newsPageFunc(vm *otto.Otto, name string) {
	obj, _ := vm.Get(name)
	common := obj.Object()
	languageCodes := common.Keys()
	for _, languageCode := range languageCodes {
		data, _ := common.Get(languageCode)
		keys := data.Object().Keys()
		for _, key := range keys {
			val, _ := data.Object().Get(key)
			kk := val.Object().Keys()
			for _, k := range kk {
				v, _ := val.Object().Get(k)
				newKey := fmt.Sprintf("%s.%s.%s", name, key, k)
				insertData(languageCode, newKey, v.String())
			}
		}
	}
}

func newsPageInfoFunc(vm *otto.Otto, name string) {
	obj, _ := vm.Get(name)
	common := obj.Object()
	languageCodes := common.Keys()
	for _, languageCode := range languageCodes {
		result := getSubObj(name, common, languageCode)
		for key, val := range result {
			insertData(languageCode, key, val.(string))
		}
	}
}

func bonusActivitiesFunc(vm *otto.Otto, name string) {
	obj, _ := vm.Get(name)
	common := obj.Object()
	languageCodes := common.Keys()
	for _, languageCode := range languageCodes {
		result := getSubObj(name, common, languageCode)
		for key, val := range result {
			insertData(languageCode, key, val.(string))
		}
	}
}

func paymentInfoFunc(vm *otto.Otto, name string) {
	obj, _ := vm.Get(name)
	common := obj.Object()
	languageCodes := common.Keys()
	for _, languageCode := range languageCodes {
		result := getSubObj(name, common, languageCode)
		for key, val := range result {
			insertData(languageCode, key, val.(string))
		}
	}
}

func termsOfServiceFunc(vm *otto.Otto, name string) {
	obj, _ := vm.Get(name)
	common := obj.Object()
	languageCodes := common.Keys()
	for _, languageCode := range languageCodes {
		result := getSubObj(name, common, languageCode)
		for key, val := range result {
			insertData(languageCode, key, val.(string))
		}
	}
}

func getSubObj(name string, object *otto.Object, languageCode string) map[string]interface{} {
	data, _ := object.Get(languageCode)
	keys := data.Object().Keys()
	result := cmap.New().Items()
	for _, key := range keys {
		val, _ := data.Object().Get(key)
		result = subSub(result, val, fmt.Sprintf("%s.%s", name, key))
	}
	return result
}

func subSub(result map[string]interface{}, obj otto.Value, key string) map[string]interface{} {
	if obj.IsObject() {
		kk := obj.Object().Keys()
		for _, subKey := range kk {
			v, _ := obj.Object().Get(subKey)
			key_ := fmt.Sprintf("%s.%s", key, subKey)
			if v.IsObject() {
				result = subSub(result, v, key_)
			} else {
				result[key_] = v.String()
			}
		}
	} else {
		result[key] = obj.String()
	}
	return result
}

func insertData(languageCode, key, val string) {
	DB := mysql.DB.Debug()
	var result models.LanguageShowApp
	err := DB.Model(models.LanguageShowApp{}).
		Where("language_code", languageCode).
		Where("type", "h5").
		Where("show_column", key).
		Find(&result).Error
	if err != nil {
		logger.Error(err)
		return
	}
	if result.Id > 0 {
		err = DB.Model(models.LanguageShowApp{}).
			Where("language_code", languageCode).
			Where("type", "h5").
			Where("show_column", key).
			Update("show_word", val).Error
		if err != nil {
			logger.Error(err)
			return
		}
	} else {
		result.LanguageCode = languageCode
		result.ShowColumn = key
		result.ShowWord = val
		result.AppPage = ""
		result.Type = "h5"
		result.Sort = 0
		result.Updater = "yezhiming"
		result.UpdateDate = time.Now()
		err = DB.Model(models.LanguageShowApp{}).Create(&result).Error
		if err != nil {
			logger.Error(err)
			return
		}
	}
}
