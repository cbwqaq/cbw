package main

import (
	"GINCHAT/driver"
	"GINCHAT/pkg/setting"
	"GINCHAT/router"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lstack-org/go-web-framework/pkg/shutdown"
	"k8s.io/klog/v2"
	"net/http"
	"time"
)

func main() {
	time.Sleep(time.Second)
	klog.InitFlags(nil)
	flag.Parse()

	gin.SetMode(setting.HTTPSetting.RunMode)

	err := driver.InitDataSource()
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", setting.HTTPSetting.HTTPPort),
		Handler:        router.InitRouter(),
		ReadTimeout:    setting.HTTPSetting.ReadTimeout,
		WriteTimeout:   setting.HTTPSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			klog.Error(err)
		}
	}()

	shutdown.NewHook().Close(
		func() {
			klog.V(1).Info("start to shutdown http server...")
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				klog.Errorf("server shutdown failed: %v", err)
			}
		},
		func() {
			klog.V(1).Info("start to disconnect mongo...")
			context.WithTimeout(context.Background(), time.Second*10)
		},
	)
}
