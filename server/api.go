package server

import (
	"encoding/json"
	"fmt"
	"github.com/dubrovin/gotest/cache"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"os"
	"path/filepath"
)

// API -
type API struct {
	router *routing.Router
	cache  *cache.Cache
}

// NewAPI -
func NewAPI(cache *cache.Cache) *API {
	return &API{
		router: routing.New(),
		cache:  cache,
	}
}

// Index is the index handler
func Index(ctx *routing.Context) error {
	fmt.Fprint(ctx, "Welcome!\n")
	return nil
}

// RegisterHandlers -
func (api *API) RegisterHandlers() {

	api.router.Get("/", Index)
	apiGroup := api.router.Group("/api")
	apiGroup.Get("/file/<name>", api.GetFile)
	apiGroup.Get("/list", api.List)
	apiGroup.Post("/upload", api.Upload)

}

// GetFile is the index handler
func (api *API) GetFile(ctx *routing.Context) error {
	fmt.Fprint(ctx, "Welcome!\n")
	return nil
}

// List is the index handler
func (api *API) List(ctx *routing.Context) error {
	ctx.SetContentType("application/json")
	list := api.cache.GetAllZipFiles()
	jsonData, err := json.Marshal(list)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		log.Println("List error ", err)
		return err
	}

	// пишем данные в ответ
	err = ctx.WriteData(jsonData)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		log.Println("List error ", err)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

//
func (api *API) Upload(ctx *routing.Context) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	for _, file := range form.File["data"] {
		out, err := os.Create(filepath.Join(api.cache.Storage.RootDir, file.Filename))
		if err != nil {
			return err
		}
		defer out.Close()

		f, err := file.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(out, f)
		if err != nil {
			return err
		}
		api.cache.AddZipFile(filepath.Join(api.cache.Storage.RootDir, file.Filename))
	}

	return nil
}
