package bcgame

import (
	"fmt"
	"testing"
)

func TestBcGameService_CreateUserToken(t *testing.T) {
	service := NewService()
	token := service.CreateUserToken("45B764D7A32", "best111", "https://fileupload.ssmartoption.com/file/download/0a3d5de4-2eb2-4001-93ea-f7c051ec3e5e.png")
	service.PushUserBalanceChange(BCPushUserBalanceChangeParams{
		PlayerId:           "",
		CurrencyName:       "",
		ChangeAmount:       0,
		BeforeChangeAmount: 0,
	})
	url := fmt.Sprintf("https://provider.dogcrash.app/api/merchants/auth/?code=%s&path=%s", token, "crash")
	fmt.Println(url)
}
