package v3

type BaseController struct {
}

type Group struct {
	BaseController
	KlineController
	NoticeController
	AgentController
	ConfigController
}
