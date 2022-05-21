package pprof

import (
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

// 开启 pprof 分析

func Debug(addr string) {
	// pprof 分析
	if conf.GetBool("app.pprof") {
		runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
		runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
		go func() {
			// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
			if err := http.ListenAndServe(addr, nil); err != nil {
				logger.Error(err)
			}
		}()
	}
}
