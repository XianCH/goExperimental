package server

import (
	"net/http"
	"time"

	"github.com/x14n/goExperimental/gin-crud-api/v3/global"
)

func SCore() {
	router := NewRouter()

	s := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if nil != err {
		global.GLogger.Info("服务启动失败" + err.Error())
	}
}
