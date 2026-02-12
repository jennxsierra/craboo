package main

import (
  "encoding/json"
  "net/http"
)

func (a *applicationDependencies)writeJSON(w http.ResponseWriter,
                                           status int, data any,
                                           headers http.Header) error  {
    jsResponse, err := json.Marshal(data)
    if err != nil {
        return err
    }
    jsResponse = append(jsResponse, '\n')
    // additional headers to be set
    for key, value := range headers {
        w.Header()[key] = value
        //w.Header().Set(key, value[0])
    }
    // set content type header
    w.Header().Set("Content-Type", "application/json")
    // explicitly set the response status code
    w.WriteHeader(status) 
    w.Write(jsResponse)

    return nil
}
