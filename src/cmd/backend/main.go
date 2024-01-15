package main

import (
	"os"
	"strconv"

	"github.com/EZCampusDevs/firepit/handler"
	"github.com/EZCampusDevs/firepit/handler/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func initLogging(app *echo.Echo) {

	IS_DEBUG := os.Getenv("DEBUG")
	LOG_LEVEL := os.Getenv("LOG_LEVEL")

	log.SetHeader("${time_rfc3339} ${level}")
	log.SetLevel(log.INFO)

	app.Logger.SetHeader("${time_rfc3339} ${level}")

	if level, err := strconv.ParseUint(LOG_LEVEL, 10, 8); err == nil {
		app.Logger.SetLevel(log.Lvl(level))
		log.Info("Read LOG_LEVEL from env: ", level)
	} else {
		log.Warn("Could not read LOG_LEVEL from env. Log level is: ", app.Logger.Level())
	}

	if debug_, err := strconv.ParseBool(IS_DEBUG); err == nil {

		app.Debug = debug_

		if debug_ {
			app.Logger.SetLevel(log.DEBUG)
		}

		log.Info("Read DEBUG from Env: ", debug_)
	} else {
		app.Debug = false
		log.Warn("Could not read DEBUG from env. Running in release mode.")
	}

	log.SetLevel(app.Logger.Level())
}

func main() {

	var e *echo.Echo
	var m *websocket.Manager

	e = echo.New()
	m = websocket.NewManager()

	initLogging(e)

	e.Use(middleware.Recover())

	e.GET("/ws", m.ServeWebsocket)
	e.GET("/room/new", m.GetRoomManager().CreateRoomGET)
	e.RouteNotFound("/", handler.Heartbeat)

	e.Logger.Fatal(e.Start(":3000"))
}