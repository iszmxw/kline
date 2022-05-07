package wss

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"goapi/pkg/logger"
	"net/http"
	"sync"
)

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsConn 声明并发安全的ws
type WsConn struct {
	*websocket.Conn              // 基础连接
	Mux             sync.RWMutex // 锁
	IsClose         bool         // 是否关闭
}

func GetSocket(c *gin.Context) (*WsConn, error) {
	if c.IsWebsocket() {
		//升级get请求为webSocket协议
		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return &WsConn{
			ws,
			sync.RWMutex{},
			false,
		}, nil
	} else {
		return nil, errors.New("不是socket請求")
	}
}

// 关闭socket资源

func WsClose(ws *WsConn) {
	if ws.IsClose {
		logger.Info("当前socket已关闭")
	} else {
		ws.Mux.Lock()
		closeErr := ws.Conn.Close()
		if closeErr != nil {
			logger.Error(closeErr)
		} else {
			ws.IsClose = true
		}
		ws.Mux.Unlock()
		logger.Info("关闭成功")
	}
}
