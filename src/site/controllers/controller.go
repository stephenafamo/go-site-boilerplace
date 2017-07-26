package controller

import (
	"github.com/gorilla/mux"
	"github.com/stephenafamo/site/config"
	"net/http"
	"reflect"
)

type Controller struct {
	Registry map[string]reflect.Type
}

var controllers = make(map[string]reflect.Type)

func init() {
	controllers["IndexController"] = reflect.TypeOf(IndexController{})
	controllers["AssetController"] = reflect.TypeOf(AssetController{})
}

func (c *Controller) Render(w http.ResponseWriter, templateName string, p interface{}) {
	if err := config.Templates.ExecuteTemplate(w, templateName, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Controller) GetVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func Get(name string) reflect.Type {
	return controllers[name]
}
