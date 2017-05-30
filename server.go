package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"github.com/go-martini/martini"
	"github.com/ryanuber/go-filecache"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"errors"
)

const cacheTime = 500;

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/cep/:id", func(params martini.Params, writer http.ResponseWriter) string {
		writer.Header().Set("Content-Type", "application/json")
		return getCep(params["id"])
	})

	m.Run()
}

func getCep(id string) string {
	cached := getFromCache(id)
	if cached != "" {
		return cached
	}
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://viacep.com.br/ws/"+id+"/json/")

	result := " "

	// make a callback function
	fooTest := func(buf []byte, userdata interface{}) bool {
		result = string(buf)

		return true
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return saveOnCache(id, result)
}

func getFromCache(id string) string {
	updater := func(path string) error {
		return errors.New("expired")
	}

	fc := filecache.New(os.TempDir()+"/cep"+id, cacheTime*time.Second, updater)

	fh, err := fc.Get()
	if err != nil {
		return ""
	}

	content, err := ioutil.ReadAll(fh)
	if err != nil {
		return ""
	}

	return string(content)
}

func saveOnCache(id string, content string) string {
	updater := func(path string) error {
		println("updater")
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write([]byte(content))
		return err
	}

	fc := filecache.New(os.TempDir()+"/cep"+id, cacheTime*time.Second, updater)

	_, err := fc.Get()
	if err != nil {
		return ""
	}

	return content
}
