package main

import (
	"boilerplate/config"
	"boilerplate/controller"
	"boilerplate/exception"
	"boilerplate/helper"
	"boilerplate/repository"
	"boilerplate/service"
	"flag"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/joho/godotenv"

	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := godotenv.Load()
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
}

func main() {
	user := flag.Bool("user", false, "display user insert")
	flag.Parse()

	mongo := config.MongoConnection()
	//mysql := config.MysqlConnection()
	if *user {
		helper.InsertUser(mongo)
	}
	//redis := config.RedisConnection()

	// Implement Repositories Here
	userRepository := repository.NewUserRepository(mongo)

	// Implement Services Here
	userService := service.NewUserService(userRepository)

	// Implement Controllers Here
	authController := controller.NewAuthController(userService)

	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: exception.ErrorHandler,
		BodyLimit:    5 * 1024 * 1024,
	})
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(pprof.New())

	// Prefix
	prefix := app.Group("/api/v1")

	// Implement Group Routes Here
	authController.Route(prefix)

	// Static Path {IMAGE_BASE_URL/images}
	app.Static("/images", "./uploads/images")
	app.Static("/csv", "./uploads/csv", fiber.Static{
		Next: func(c *fiber.Ctx) bool {
			c.Attachment("nps_export.csv")
			return false
		},
	})

	err := app.Listen(":" + os.Getenv("PORT"))
	exception.PanicIfNeeded(err)
}
