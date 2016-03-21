package main

import (
  "fmt"
  "net/http"
  "os"
  "log"
  "github.com/SlyMarbo/rss"
  "encoding/json"
  "strconv"
)

func introHandler(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(res, "rss-machine")
}

func parseHandler(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-Type", "application/json; charset=utf-8")
  query := req.URL.Query()
  url := query.Get("url")

  rss.CacheParsedItemIDs(false)
  feed, error := rss.Fetch(url)
  if error != nil {
    json, _ := json.Marshal(map[string]interface{}{"error": error.Error()})
    http.Error(res, string(json), http.StatusInternalServerError)
    errorsLogger.Println(error)
    return
  }

  body, error := json.Marshal(feed)
  if error != nil {
    json, _ := json.Marshal(map[string]interface{}{"error": error.Error()})
    http.Error(res, string(json), http.StatusInternalServerError)
    errorsLogger.Println(error)
    return
  }

  res.Write(body)
}

var logger *log.Logger
var errorsLogger *log.Logger

func main() {
  logger = log.New(os.Stdout, "rss-macine ", log.Ldate | log.Ltime | log.Lshortfile)
  errorsLogger = log.New(os.Stderr, "rss-macine ", log.Ldate | log.Ltime | log.Lshortfile)

  http.HandleFunc("/", introHandler)
  http.HandleFunc("/parse", parseHandler)

  portAsString := os.Getenv("PORT")
  port := 5000
  if portAsString != "" {
    port, _ = strconv.Atoi(portAsString)
  }

  logger.Println("Started listening on", port)
  err := http.ListenAndServe(":" + strconv.Itoa(port), nil)
  if err != nil {
    errorsLogger.Panic(err)
  }
}
