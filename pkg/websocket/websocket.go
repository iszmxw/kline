package wss

import (
	"container/list"
	"errors"
	"goapi/app/models"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	EnableCompression: false, // 苹果某些旧版本的ios不兼容该压缩功能，所以苹果机型直接禁用掉就好了，或者找出相关兼容的机型进行白名单
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsConn 声明并发安全的ws
type WsConn struct {
	*websocket.Conn              // 基础连接
	Mux             sync.RWMutex // 锁
	IsClose         bool         // 是否关闭
	AppUserId       string       // 用户 ID
	OnceNotice      sync.Once    // 通知方法
	SecWebSocketKey string       // SecWebSocketKey
	Ip              string       // 用户的ip
	List            *list.List
}

func GetSocket(c *gin.Context) (*WsConn, error) {
	var ws *websocket.Conn
	var err error
	if c.IsWebsocket() {
		if c.GetBool("socket") {
			wsAny, _ := c.Get("ws")
			ws = wsAny.(*websocket.Conn)
		} else {
			c.Set("socket", true)
			//升级get请求为webSocket协议
			ws, err = upGrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				logger.Error(err)
				return nil, err
			}
			c.Set("ws", ws)
		}
		return &WsConn{
			ws,
			sync.RWMutex{},
			false,
			"",
			sync.Once{},
			"---",
			"127.0.0.1",
			list.New(),
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
		if len(ws.AppUserId) > 0 {
			redis.Client.SRem("onlineUsers", ws.AppUserId)
			mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", ws.AppUserId).Update("online", 0)
		}
		if len(ws.Ip) > 0 {
			redis.Client.SRem("onlineConnect", ws.Ip)
		}
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
