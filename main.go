package main

import (
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/zainafifaldi/digicafe/infrastructure/auth"
	"github.com/zainafifaldi/digicafe/infrastructure/persistence"
	"github.com/zainafifaldi/digicafe/interfaces"
	"github.com/zainafifaldi/digicafe/interfaces/middleware"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	dbDriver := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	parseTime := "true"

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	services, err := persistence.NewRepositories(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName, parseTime)
	if err != nil {
		panic(err)
	}
	defer services.Close()

	redisService, err := auth.NewRedisDB(redisHost, redisPort, redisPassword)
	if err != nil {
		log.Fatal(err)
	}

	tk := auth.NewToken()
	// fu := fileupload.NewFileUpload()

	stores := interfaces.NewStore(services.Store, redisService.Auth, tk)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //For CORS

	r.GET("/stores", stores.GetAllStore)
	r.GET("/stores/:id", stores.GetStoreDetail)

	//Starting the application
	app_port := os.Getenv("SERVER_PORT")
	if app_port == "" {
		log.Fatal(errors.New("Port is not defined"))
	}
	log.Fatal(r.Run(":" + app_port))
}
