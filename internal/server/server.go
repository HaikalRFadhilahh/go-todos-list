package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HaikalRFadhilahh/go-todos-list/internal/router"
	"github.com/HaikalRFadhilahh/go-todos-list/pkg/helper"
	"github.com/joho/godotenv"
)

type ApiServer struct {
	ServerListenAndServe string
}

func InitApiServer() *ApiServer {
	_ = godotenv.Load()
	return &ApiServer{
		ServerListenAndServe: fmt.Sprintf("%s:%s", helper.GetEnv("HOST", "127.0.0.1"), helper.GetEnv("PORT", "8000")),
	}
}

func (a *ApiServer) RunServer() {
	// Define Init Router
	r := router.Router()

	// Start Router
	fmt.Println("Server Starting....")
	fmt.Println("Server Running On", a.ServerListenAndServe)
	if err := http.ListenAndServe(a.ServerListenAndServe, r); err != nil {
		log.Fatalln("Server Error :", err.Error())
	}

}
