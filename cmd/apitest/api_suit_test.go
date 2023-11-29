package apitest

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/daolis/training/todo-api/cmd/api"
)

const (
	baseURL = "http://localhost:8081/api"
)

type ApiTestSuite struct {
	suite.Suite
	quit    chan os.Signal
	baseURL string
}

func (suite *ApiTestSuite) SetupSuite() {
	suite.quit = make(chan os.Signal)
	started := make(chan struct{})

	config := &api.AppConfig{
		Addr: ":8081",
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	go func() {
		app := api.NewApp(config, logger)
		err := app.StartServer()
		if err != nil {
			suite.Suite.T().FailNow()
		}
		close(started)
		<-suite.quit
	}()
}

func (suite *ApiTestSuite) TearDownSuite() {
	close(suite.quit)
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, &ApiTestSuite{
		baseURL: baseURL,
	})
}
