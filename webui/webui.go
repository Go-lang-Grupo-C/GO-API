package webui

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"

	"github.com/Go-lang-Grupo-C/GO-API/server"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//go:embed dist/*
var Assets embed.FS

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

func AssetHandler(prefix, root string) http.Handler {

	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)

		f, err := Assets.Open(assetPath)
		if os.IsNotExist(err) {
			return Assets.Open("dist/spa/index.html")
		}

		return f, err
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}

func NewWebUI() http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/api/v1").Handler(http.StripPrefix("/api/v1", server.NewServer()))

	router.PathPrefix("/").Handler(AssetHandler("/api/v1", "dist/spa/"))

	n := negroni.Classic()
	n.UseHandler(router)

	return n
}
