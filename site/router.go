package main

import (
	"github.com/gorilla/mux"
	"path/to/your/app/controllers"
	"net/http"
	"reflect"
)

func customRouter() *stephenRouter {
	muxRouter := mux.NewRouter()
	router := new(stephenRouter)
	router.Router = *muxRouter.StrictSlash(true)
	return router
}

type stephenRouter struct {
	mux.Router
}

func (c *stephenRouter) handler(path string, controllerName string, method string) {
	theController := reflect.New(controller.Get(controllerName))
	theMethod := theController.MethodByName(method)
	c.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		theMethod.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r), theController.Elem().FieldByName("Params")})
	})
}
