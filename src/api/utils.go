package api

func JsonOutput(data interface{}, err error, statusCode int) map[string]interface{} {
    output := map[string]interface{}{"data": data, "error": nil, "statusCode": statusCode}
    if err != nil {
        output["error"] = map[string]interface{}{"message": err.Error()}
    }
    return output
}
