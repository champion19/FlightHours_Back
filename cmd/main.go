package main

import (
"github.com/gin-gonic/gin"
"github.com/champion19/Flighthours_backend/server"
"log/slog"
)
func main(){
	gin.SetMode(gin.ReleaseMode)
	app:=gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	dependencies:= server.Boostrap(app)
	serverAddr:=dependencies.Config.GetServerAddress()
	slog.Info("Starting server",slog.String("address", serverAddr))

  if err := app.Run(serverAddr); err != nil {
		slog.Error("Server failed to start", slog.String("error", err.Error()))
		return
}

}
