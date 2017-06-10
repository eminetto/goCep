package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"github.com/gorilla/mux"
	"github.com/ryanuber/go-filecache"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"errors"
)

const serverPort = "3000";
const cacheTime = 500;

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world!"))
	})
	router.HandleFunc("/cep/{id}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(getCep(vars["id"])))
	})

	http.ListenAndServe(":"+serverPort, router)
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

	fc := filecache.New(getCacheFilename(id), cacheTime*time.Second, updater)

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

	fc := filecache.New(getCacheFilename(id), cacheTime*time.Second, updater)

	_, err := fc.Get()
	if err != nil {
		return ""
	}

	return content
}

func getCacheFilename(id string) string {
  return os.TempDir()+"/cep"+id;
}
