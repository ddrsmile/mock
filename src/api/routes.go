package api

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
    "net/http"
    "os"
    "path"
    "sync"
)

type MockRouter struct {
    mu sync.Mutex
    r *mux.Router
}

func (mr *MockRouter) Swap(r *mux.Router) {
    mr.mu.Lock()
    mr.r = r
    mr.mu.Unlock()
}

func (mr *MockRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    mr.r.ServeHTTP(w, r)
}

func getItems(fileName string) ([]Item, error) {
    f, err := os.Open(path.Join(ResourceDir, fileName))
    defer f.Close()
    if err != nil {
        return nil, err
    }
    byteVar, _ := ioutil.ReadAll(f)
    var items []Item
    err = json.Unmarshal(byteVar, &items)
    if err != nil {
        return nil, err
    }
    return items, nil
}

func SetRoutes(r *mux.Router) {
    mr := &MockRouter{}
    // manage
    manage := r.PathPrefix("/mockmanage").Subrouter()
    manage.Use(SetJsonContentType)
    manage.HandleFunc("/upload", GetUploadHandler())
    manage.HandleFunc("/download", GetDownloadHandler())
    manage.HandleFunc("/reload", GetReloadHandler(mr))

    // api
    api := r.PathPrefix("/mockapi").Subrouter()
    api.Handle("/", mr)
    api.Handle("/{_:.*}", mr)
    // set api router
    apiRouter, _ := GetApiRouter(DefaultLoadFileName)
    mr.Swap(apiRouter)
}

func GetApiRouter(fileName string) (*mux.Router, error) {
    r := mux.NewRouter().PathPrefix("/mockapi").Subrouter()
    items, err := getItems(fileName)
    if err != nil {
        return nil, err
    }
    for _, item := range items {
        r.HandleFunc(item.Path, GetMockApiHandler(item)).Methods(item.Methods...)
    }
    return r, nil
}
