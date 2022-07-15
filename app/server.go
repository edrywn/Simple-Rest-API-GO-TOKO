package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edrywn/toko-online/app/item"
	"github.com/edrywn/toko-online/app/seller"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Welcome to" + appConfig.AppName)

	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/go_tokoOnline?charset=utf8mb4&parseTime=True&loc=Local"
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Filed connect to database")
	}

	server.DB.AutoMigrate(&item.Item{})
	server.DB.AutoMigrate(&seller.Seller{})

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listning to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	server := Server{}
	appConfig := AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading ENV file")
	}

	appConfig.AppName = getEnv("APP_NAME", "TokoOnline")
	appConfig.AppEnv = getEnv("APP_ENV", "Development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
