package apitest

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/daolis/training/todo-api/internal/data"
)

var todoTestTypes = []string{"APITestTodo 01", "APITestTodo 02"}

func (suite *ApiTestSuite) TestTodosCRUD() {
	testEntityIds := make([]string, len(todoTestTypes))

	suite.T().Run("Create todos", func(t *testing.T) {
		r := require.New(t)
		for idx, testEntity := range todoTestTypes {
			newEntity := new(data.Todo)
			suite.NewTextCall(t, fmt.Sprintf("Create test todo%2d", idx+1), http.MethodPost, "/toDoItems",
				WithResponseStatus(http.StatusCreated),
				WithRequestData(testEntity),
				WithResponseData(newEntity))

			r.NotNil(newEntity)
			r.NotEmpty(newEntity.ID)
			r.False(newEntity.Done)
			testEntityIds[idx] = newEntity.ID

			r.Equal(testEntity, newEntity.Description)
		}

	})
	suite.T().Run("Get todos list", func(t *testing.T) {
		r := require.New(t)
		var entities []data.Todo
		suite.NewTextCall(t, "Get todo list", http.MethodGet, "/toDoItems",
			WithResponseData(&entities))

		r.NotNil(entities)
		r.Len(entities, 4)
		r.Equal("test todo 01", entities[0].Description)
		r.Equal("test todo 02", entities[1].Description)
		r.Equal("APITestTodo 01", entities[2].Description)
		r.Equal("APITestTodo 02", entities[3].Description)
	})
	suite.T().Run("Get todo by id", func(t *testing.T) {
		r := require.New(t)
		var entity data.Todo
		scndTestEntryId := testEntityIds[1]
		suite.NewTextCall(t, "Get todo 1", http.MethodGet, fmt.Sprintf("/toDoItems/%s", scndTestEntryId),
			WithResponseData(&entity))

		r.NotNil(entity)
		r.Equal(scndTestEntryId, entity.ID)
		r.Equal("APITestTodo 02", entity.Description)
		r.Equal(false, entity.Done)
	})
	suite.T().Run("Get todo by id (not exists)", func(t *testing.T) {
		suite.NewTextCall(t, "Get todo 1", http.MethodGet, "/toDoItems/aaaaaaaaaaa",
			WithResponseStatus(http.StatusNotFound))
	})
	suite.T().Run("Set todo to done", func(t *testing.T) {
		r := require.New(t)
		var entity data.Todo
		scndTestEntryId := testEntityIds[1]
		suite.NewTextCall(t, "Set todo to done", http.MethodPost, fmt.Sprintf("/toDoItems/%s/setDone", scndTestEntryId),
			WithResponseData(&entity))

		r.NotNil(entity)
		r.Equal(scndTestEntryId, entity.ID)
		r.Equal("APITestTodo 02", entity.Description)
		r.Equal(true, entity.Done)
	})
	suite.T().Run("Get todos list (updated)", func(t *testing.T) {
		r := require.New(t)
		var entities []data.Todo
		suite.NewTextCall(t, "Get todo list", http.MethodGet, "/toDoItems",
			WithResponseData(&entities))

		r.NotNil(entities)
		r.Len(entities, 4)
		r.Equal("test todo 01", entities[0].Description)
		r.Equal(false, entities[0].Done)
		r.Equal("test todo 02", entities[1].Description)
		r.Equal(false, entities[1].Done)
		r.Equal("APITestTodo 01", entities[2].Description)
		r.Equal(false, entities[2].Done)
		r.Equal("APITestTodo 02", entities[3].Description)
		r.Equal(true, entities[3].Done)
	})
}
