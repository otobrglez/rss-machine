package main

import (
  "testing"
  "os"
  "flag"
  "net/http"
  "net/http/httptest"
  "log"
)

func init() {
  logger = log.New(os.Stdout, "rss-macine ", log.Ldate | log.Ltime | log.Lshortfile)
  errorsLogger = log.New(os.Stderr, "rss-macine ", log.Ldate | log.Ltime | log.Lshortfile)
}

func TestMain(m *testing.M) {
  flag.Parse()
  os.Exit(m.Run())
}

func TestSuccessfulParse(t *testing.T) {
  url := "http://rss.cnn.com/rss/edition.rss"
  req, _ := http.NewRequest("GET", "/parse?url=" + url, nil)
  w := httptest.NewRecorder()

  parseHandler(w, req)
  if w.Code != http.StatusOK {
    t.Errorf("Could not parse")
  }
}

func TestEmptyParse(t *testing.T) {
  req, _ := http.NewRequest("GET", "/parse?url=", nil)
  w := httptest.NewRecorder()

  parseHandler(w, req)
  if w.Code != http.StatusBadRequest {
    t.Errorf("Should be StatusBadRequest")
  }
}

func TestFailedParse(t *testing.T) {
  url := "http://rss.cnn.com/rss/editionx.rss"
  req, _ := http.NewRequest("GET", "/parse?url=" + url, nil)
  w := httptest.NewRecorder()

  parseHandler(w, req)
  if w.Code != http.StatusInternalServerError {
    t.Errorf("Should return StatusInternalServerError")
  }
}
