package api

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path"
    "time"
)

const (
    ResourceDir = "../resources"
    DefaultDownloadFileName = "default.json"
    DefaultUploadFileName = "uploaded.json"
)

// mock manage
func GetUploadHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data := map[string]interface{}{"status": "json data uploaded."}
        statusCode := http.StatusOK
        fileName := r.FormValue("fileName")
        if fileName == "" {
            fileName = DefaultUploadFileName
        }
        byteVar, err := ioutil.ReadAll(r.Body)
        if err != nil {
            data["status"] = "fail to upload json data."
            statusCode = http.StatusBadRequest
        } else {
            _ = ioutil.WriteFile(path.Join(ResourceDir, fileName), byteVar, 0644)
        }
        w.WriteHeader(statusCode)
        _ = json.NewEncoder(w).Encode(JsonOutput(data, err, statusCode))
    }
}

func GetDownloadHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        statusCode := http.StatusOK
        fileName := r.FormValue("fileName")
        if fileName == "" {
            fileName = DefaultDownloadFileName
        }
        f, err := os.Open(path.Join(ResourceDir, fileName))
        defer f.Close()
        if err != nil {
            data := map[string]interface{}{"status": "fail to download json data."}
            statusCode = http.StatusBadRequest
            _ = json.NewEncoder(w).Encode(JsonOutput(data, err, statusCode))
        } else {
            w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%v", fileName))
            http.ServeContent(w, r, fileName, time.Now(), f)
        }
    }
}

func GetReloadHandler(mockRouter *MockRouter) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data := map[string]interface{}{"status": "mock api router is reload."}
        statusCode := http.StatusOK
        router, err := GetApiRouter()
        if err != nil {
            data["status"] = "fail to reload mock api router"
            statusCode = http.StatusBadRequest
        } else {
            mockRouter.Swap(router)
        }
        w.WriteHeader(statusCode)
        output := JsonOutput(data, err, statusCode)
        _ = json.NewEncoder(w).Encode(output)
    }
}

// mock api
func GetMockApiHandler(item Item) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", item.Response.ContentType)

        for k, v := range item.Response.Headers {
            w.Header()[k] = v
        }

        for _, cookie := range item.Response.Cookies {
            http.SetCookie(w, &cookie)
        }

        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte(item.Response.Content))
    }
}
