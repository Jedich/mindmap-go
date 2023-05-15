package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"mindmap-go/app/controllers"
	"mindmap-go/app/repository"
	"mindmap-go/app/services"
	"mindmap-go/internal/database"
	"mindmap-go/router"
	"mindmap-go/utils/response"
	"net/http"
	"os"
	"testing"
)

func initDB(t testing.TB) *database.Database {
	conn, err := gorm.Open(mysql.Open("root:11@tcp(localhost:3306)/mindmap_test?parseTime=true"), &gorm.Config{})
	assert.NoError(t, err)

	db = &database.Database{Connection: conn}
	err = db.Migrate()
	assert.NoError(t, err)

	repo = repository.NewCardRepository(db)
	return &database.Database{Connection: conn}
}

func initApp(conn *database.Database) *fiber.App {
	userRepo := repository.NewUserRepository(conn)
	accRepo := repository.NewAccountRepository(conn)
	service := services.NewUserService(userRepo, accRepo)
	controller := controllers.NewController(service,
		services.NewMapService(repository.NewMapRepository(conn)),
		services.NewCardService(repository.NewCardRepository(conn)))
	app := fiber.New(fiber.Config{
		ErrorHandler: response.ErrorHandler,
	})
	router.NewUserRouter(app, controller).RegisterAuthRoutes()
	return app
}

func TestAll(t *testing.T) {
	conn := initDB(t)
	defer func() {
		err := db.Connection.Exec("DROP TABLE accounts, cards, files, maps, users;").Error
		if err != nil {
			panic(err)
		}
		db.CloseConnection()
	}()
	err := os.Setenv("JWT_SECRET", "test")
	assert.NoError(t, err)
	app := initApp(conn)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(app)).
		Post("/auth/register").
		JSON(`{"username": "a"}`).
		Expect(t).
		Status(http.StatusForbidden).
		End()

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(app)).
		Post("/auth/register").
		Expect(t).
		Status(http.StatusUnprocessableEntity).
		End()

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(app)).
		Post("/auth/register").
		JSON(`{"username": "aaaaa", "email": "a@b.c", "password": "11111111"}`).
		Expect(t).
		CookiePresent("token").
		Status(http.StatusOK).
		End()

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(app)).
		Post("/auth/login").
		JSON(`{"email": "a"}`).
		Expect(t).
		Status(http.StatusForbidden).
		End()

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(app)).
		Post("/auth/login").
		Expect(t).
		Status(http.StatusUnprocessableEntity).
		End()

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(app)).
		Post("/auth/login").
		JSON(`{"email": "a@b.c", "password": "11111111"}`).
		Expect(t).
		CookiePresent("token").
		Status(http.StatusOK).
		End()
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		// copy body
		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}
