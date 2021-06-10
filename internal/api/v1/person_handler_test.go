package v1

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/disturb16/go-clean-architecture/dbutils"
	"github.com/disturb16/go-clean-architecture/internal/persons/repository"
	"github.com/disturb16/go-clean-architecture/internal/persons/service"
	"github.com/disturb16/go-clean-architecture/settings"
	"github.com/labstack/echo/v4"
	"github.com/sanservices/apicore/validator"
)

var h *Handler

func TestMain(m *testing.M) {

	ctx := context.Background()
	config, err := settings.Get("../../settings")

	if err != nil {
		log.Panic(err)
	}

	// get db connection
	db, err := dbutils.CreateSqliteConnection(config)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	repo, err := repository.New(ctx, config, db)
	if err != nil {
		log.Panic(err)
	}

	// Setup dependencies
	srv := service.New(repo)
	v := validator.NewValidator()
	h = NewHandler(config, srv, v)

	code := m.Run()

	// Cleanup database from the project
	os.Remove(config.DB.Name + ".db")
	os.Exit(code)
}

func TestPersons(t *testing.T) {
	testCases := []struct {
		Name               string
		Params             map[string]string
		StatusCodeExpected int
	}{
		{
			Name:               "Should retrieve data",
			Params:             map[string]string{"limit": "1"},
			StatusCodeExpected: http.StatusOK,
		},
	}

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel() // This allows tests to run in parallel mode

			e := echo.New()
			c := e.NewContext(r, w)

			// set url query parameters
			q := r.URL.Query()
			for k, v := range tc.Params {
				q.Add(k, v)
			}

			r.URL.RawQuery = q.Encode()
			h.persons(c)

			if w.Code != tc.StatusCodeExpected {
				t.Errorf("Expected statusCode %d instead received: %d", tc.StatusCodeExpected, w.Code)
			}
		})
	}
}
