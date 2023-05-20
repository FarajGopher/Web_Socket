package main

import (
	"auth_web_sockets/db"
	"auth_web_sockets/internal/router"
	"auth_web_sockets/internal/user"
	"log"
)

func main() {
	dbConn, err := db.InitDatabase()
	if err != nil {
		log.Println("initializing database error:", err)
		return
	}
	log.Println("database successfully connected....")
	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
