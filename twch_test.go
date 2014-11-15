package twch

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "net/url"
)

var (
  mux    *http.ServeMux
  server *httptest.Server
  client *Client
)

func setup() {
  mux = http.NewServeMux()
  server = httptest.NewServer(mux)
  client, _ = NewClient("test-client-key")
  client.BaseUrl, _ = url.Parse(server.URL)
  fmt.Printf("TWCH: %+v", client)
}

func teardown() {
  server.Close()
}
