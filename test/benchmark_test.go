package test

import (
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func BenchmarkLogin(b *testing.B) {
	conn := initDB(b)
	defer func() {
		err := db.Connection.Exec("DROP TABLE accounts, cards, files, maps, users;").Error
		if err != nil {
			panic(err)
		}
		db.CloseConnection()
	}()
	err := os.Setenv("JWT_SECRET", "test")
	assert.NoError(b, err)
	app := initApp(conn)

	for i := 0; i < b.N; i++ {
		apitest.New().
			HandlerFunc(FiberToHandlerFunc(app)).
			Post("/auth/login").
			JSON(fmt.Sprintf(`{"email": "a%d@b.c", "password": "11111111"}`, i)).
			Expect(b).
			Status(http.StatusUnauthorized).
			End()
	}
}
