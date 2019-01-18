package handler

import (
	"github.com/golang/glog"
	"github.com/k8sinstance/go/jutils"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

//import (
//	"gopkg.in/mgo.v2"
//)

//type (
// Handler - Handle with DB connection
//	Handler struct {
//			DB *mgo.Session
//			}
//		)

type Handler struct {
	//	DB *Mongo
	CON *int
}

func (h *Handler) GetHostName(c echo.Context) error {
	hn := jutils.GetHostName()
	glog.Info(hn)
	return c.JSON(http.StatusCreated, hn)
}

//curl -H "SLEEPTIME: 3"  http://localhost:8080/sleep  -v
func (h *Handler) GetToSleep(c echo.Context) error {
	//sleeptime := c.Request().Header.Get("SLEEPTIME")
	sleeptime := c.Request().Header.Get("SLEEPTIME")
	ist, _ := strconv.Atoi(sleeptime)
	glog.Info(sleeptime)
	time.Sleep(time.Duration(ist) * time.Second)

	return c.JSON(http.StatusOK, sleeptime)
}

func (h *Handler) GetToSleepInf(c echo.Context) error {
	//sleeptime := c.Request().Header.Get("SLEEPTIME")
	sleeptime := "10000000"
	ist, _ := strconv.Atoi(sleeptime)
	glog.Info(sleeptime)
	time.Sleep(time.Duration(ist) * time.Second)

	return c.JSON(http.StatusOK, sleeptime)
}
