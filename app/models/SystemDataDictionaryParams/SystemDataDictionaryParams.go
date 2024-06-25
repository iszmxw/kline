package SystemDataDictionaryParams

// Name 枚举类型
type Name int

const (
	APK_DOWNLOAD_ADDRESS Name = iota
	APK_UPLOAD_ADDRESS
	APK_LAST_UPLOAD
	ASSETS_INSERT_FRAME_CONFIG
	SMT_APK_DOWNLOAD_ADDRESS
	SMT_APK_LAST_UPLOAD
	IMAGES_UPLOAD_ADDRESS
	IMAGES_UPLOAD_UPDATE_ADDRESS
	VEST
)

// SystemDataDictionaryParams 存储键值对应关系和备注信息的结构体
type SystemDataDictionaryParams struct {
	Name   string // 键值
	Remark string // 备注信息
}

// Map 存储键值对应关系和备注信息的映射
var Map = map[Name]SystemDataDictionaryParams{
	APK_DOWNLOAD_ADDRESS:         {"SMO_APK_DOWNLOAD_ADDRESS", "安卓包下载地址"},
	APK_UPLOAD_ADDRESS:           {"SMO_APK_UPLOAD_ADDRESS", "安卓包文件上传地址"},
	APK_LAST_UPLOAD:              {"SMO_APK_LAST_UPLOAD_NAME", "安卓包最后一次更新信息"},
	ASSETS_INSERT_FRAME_CONFIG:   {"SMO_ASSETS_INSERT_FRAME", "资产插帧配置"},
	SMT_APK_DOWNLOAD_ADDRESS:     {"SMT_APK_DOWNLOAD_ADDRESS", "SMT马甲包下载地址"},
	SMT_APK_LAST_UPLOAD:          {"SMT_APK_LAST_UPLOAD_NAME", "SMT马甲包最后一次上传信息"},
	IMAGES_UPLOAD_ADDRESS:        {"IMAGES_UPLOAD_ADDRESS", "base64 图片上传地址"},
	IMAGES_UPLOAD_UPDATE_ADDRESS: {"IMAGES_UPLOAD_UPDATE_ADDRESS", "base64 图片更新地址"},
	VEST:                         {"VEST", "马甲包信息参数获取"},
}

// Info 根据键代码获取键值
func Info(name Name) SystemDataDictionaryParams {
	return Map[name]
}
