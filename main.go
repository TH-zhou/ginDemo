package main

import (
	"fmt"
	"ginDemo02/pkg/setting"
	"ginDemo02/routers"
	"net/http"
)

func main() {
	engine := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        engine,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
