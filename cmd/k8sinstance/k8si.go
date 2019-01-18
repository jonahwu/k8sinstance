package main

import (
	"flag"
	//"fmt"
	"github.com/golang/glog"
	"github.com/k8sinstance/cmd/handler"
	"github.com/k8sinstance/go/jutils"
	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

func main() {
	flag.Parse()
	// here we setup the default stdout for glog
	flag.Lookup("logtostderr").Value.Set("true")
	data := jutils.GetHello()
	glog.Info(data)
	//fmt.Println(data)
	glog.Info("now into data")
	//glog.Error("glog in errorla")

	//db, err := mgo.Dial("mongo")
	//if err != nil {
	//		e.Logger.Fatal(err)
	//	}
	//	h := &handler.Handler{DB: db}
	//	h := &Handler
	//	var h Handler
	// replace the following and by using db connection here
	// and send the pointer to handler
	//see https://github.com/petronetto/echo-mongo-api/blob/master/handler/handler.go
	con := 5
	h := &handler.Handler{CON: &con}

	e := echo.New()

	e.GET("/*", h.GetHostName)
	e.GET("/sleep", h.GetToSleep)
	e.GET("/sleepinf", h.GetToSleepInf)
	e.Logger.Fatal(e.Start(":8000"))

	glog.Flush()
	select {}
}
