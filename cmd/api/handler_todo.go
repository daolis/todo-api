package api

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/daolis/training/todo-api/internal/data"
)

func (a *app) listTodos(writer http.ResponseWriter, request *http.Request) {
	todos := data.Todos.GetTodos()
	err := a.writeJSON(writer, http.StatusOK, todos)
	if err != nil {
		a.logger.Error(err.Error())
	}
}

func (a *app) addTodo(writer http.ResponseWriter, request *http.Request) {
	content, err := io.ReadAll(request.Body)
	if err != nil {
		a.errorResponse(writer, request, http.StatusInternalServerError, "could not read body")
	}

	todo := data.Todos.CreateTodo(string(content))
	err = a.writeJSON(writer, http.StatusCreated, todo)
	if err != nil {
		a.logger.Error(err.Error())
	}
}

func (a *app) getTodo(writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id := params.ByName("id")
	if id == "" {
		a.notFoundResponse(writer, request)
		return
	}
	todo, err := data.Todos.GetTodoById(id)
	if err != nil {
		a.notFoundResponse(writer, request)
		return
	}
	err = a.writeJSON(writer, http.StatusOK, todo)
	if err != nil {
		a.logger.Error(err.Error())
	}
}

func (a *app) setTodoDone(writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id := params.ByName("id")
	if id == "" {
		a.notFoundResponse(writer, request)
		return
	}
	todo, err := data.Todos.SetDone(id)
	if err != nil {
		a.notFoundResponse(writer, request)
		return
	}
	err = a.writeJSON(writer, http.StatusOK, todo)
	if err != nil {
		a.logger.Error(err.Error())
	}
}
