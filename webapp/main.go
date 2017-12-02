package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/eaciit/clit"
	"github.com/eaciit/knot/knot.v1"

	"eaciit/tplmaterial/model"
	"eaciit/tplmaterial/webapp/controller"
)

var (
	err error
)

func main() {
	clit.LoadConfigFromFlag("", "", filepath.Join(clit.ExeDir(), "..", "config", "app.json"))
	if err = clit.Commit(); err != nil {
		kill(err)
	}
	defer clit.Close()

	if err = model.SetConnectionString(clit.Config("default", "DbConnection", "").(string)); err != nil {
		kill(err)
	}

	app := createApp()
	webhost := clit.Config("default", "WebHost", "").(string)

	routes := map[string]knot.FnContent{
		"/": func(r *knot.WebContext) interface{} {
			http.Redirect(r.Writer, r.Request, "/home/default", http.StatusTemporaryRedirect)
			return true
		},
		/*
			"prerequest": func(r *knot.WebContext) interface{} {
				url := r.Request.URL.String()

				if url == "/login/default" && r.Session("username") != nil {
					http.Redirect(r.Writer, r.Request, "/dashboard/default", http.StatusTemporaryRedirect)
					return true
				}

				if strings.Index(url, "/login") < 0 && url != "/" {
					//h.IsAuthenticate(r)
					return nil
				}
				return nil
			},
			"postrequest": func(r *knot.WebContext) interface{} {
				return nil
			},
		*/
	}

	knot.StartAppWithFn(app, webhost, routes)
}

func createApp() *knot.App {
	dict := map[string]string{
		"workdir": clit.ExeDir(),
	}

	webroot := translateToStr(clit.Config("default", "WebRoot", "").(string), dict)
	if webroot == "" {
		webroot = clit.ExeDir()
	}

	app := knot.NewApp("ecapp")

	/**REGISTER ALL CONTROLLERS HERE**/
	app.Register(new(controller.Home))
	app.DefaultOutputType = knot.OutputTemplate

	/* FOLDER STRUCTURE */
	dict["workdir"] = webroot
	app.Static("asset", translateToStr(filepath.Join(webroot, "asset"), dict))
	/* END FOLDER STRUCTURE */

	app.LayoutTemplate = "_layout.html"
	app.ViewsPath = filepath.Join(webroot, "views")

	return app
}

func translateToStr(source string, dict map[string]string) string {
	ret := source
	for k, v := range dict {
		find := "{$" + k + "}"
		ret = strings.Replace(ret, find, v, -1)
	}
	return ret
}

func kill(err error) {
	fmt.Printf("error. %s \n", err.Error())
	os.Exit(100)
}
