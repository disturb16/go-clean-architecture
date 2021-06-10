package service

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/disturb16/go-clean-architecture/dbutils"
	"github.com/disturb16/go-clean-architecture/internal/persons"
	"github.com/disturb16/go-clean-architecture/internal/persons/repository"
	"github.com/disturb16/go-clean-architecture/settings"
)

var service persons.Service

func TestMain(m *testing.M) {
	ctx := context.Background()
	config, err := settings.Get("../../settings")
	if err != nil {
		log.Panic(err)
	}

	// get database conection
	db, err := dbutils.CreateSqliteConnection(config)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	repo, err := repository.New(ctx, config, db)
	if err != nil {
		log.Panic(err)
	}

	service = New(repo)

	code := m.Run()

	// Cleanup database
	os.Remove(config.DB.Name + ".db")
	os.Exit(code)
}

func TestSavePerson(t *testing.T) {
	useCases := []struct {
		Name          string
		PersonNanme   string
		PersonAge     int
		ExpectedError error
	}{
		{
			Name:          "Should insert person",
			PersonNanme:   "Someone",
			PersonAge:     20,
			ExpectedError: nil,
		},

		{
			Name:          "Should fail with empty name",
			PersonNanme:   "",
			PersonAge:     20,
			ExpectedError: errNameEmpty,
		},

		{
			Name:          "Should fail with invalid age",
			PersonNanme:   "Greg",
			PersonAge:     15,
			ExpectedError: errAgeNotValid,
		},
	}

	ctx := context.Background()

	for i := range useCases {

		// Get the use case with the current index.
		// This is done in this way to avoid issues in parallel mode.
		uc := useCases[i]

		t.Run(uc.Name, func(t *testing.T) {
			t.Parallel() // Run the tests in parallel mode

			_, err := service.SavePerson(ctx, uc.PersonNanme, uc.PersonAge)

			if err != uc.ExpectedError {
				t.Errorf("Expected error: %v instead got: %v", uc.ExpectedError, err)
			}
		})
	}
}
