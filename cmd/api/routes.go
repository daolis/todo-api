package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (a *app) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/api/toDoItems", a.listTodos)
	router.HandlerFunc(http.MethodGet, "/api/toDoItems/:id", a.getTodo)
	router.HandlerFunc(http.MethodPost, "/api/toDoItems", a.addTodo)
	router.HandlerFunc(http.MethodPost, "/api/toDoItems/:id/setDone", a.setTodoDone)

	c := alice.New()
	c = c.Append(a.recoverPanic, a.enableCORS)
	chain := c.Then(router)
	return chain

}
