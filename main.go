/**
 * @Author: yin
 * @Description:main
 * @Version: 1.0.0
 * @Time : 2019-09-01 19:45
 */
package main

import (
	"fmt"
	"gin-blog/routers"
	"net/http"

	"gin-blog/pkg/setting"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
