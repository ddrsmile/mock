package api

import (
    "encoding/json"
    "net/http"
)

// mock manage
func GetManageReloadHandler(mockRouter *MockRouter) http.HandlerFunc {
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
