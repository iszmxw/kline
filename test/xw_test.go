package test

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func Test_xw(t *testing.T) {
	onlineOnlyKey := fmt.Sprintf("G%v%v", uuid.New().ID(), uuid.New().ID())
	fmt.Println(onlineOnlyKey)
}
