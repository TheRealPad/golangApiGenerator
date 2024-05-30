package integrationTests

import (
	"bytes"
	"github.com/gorilla/mux"
	"httpServer/src/controller"
	"httpServer/src/core"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/middlewares/logging"
	"httpServer/src/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startServer() *httptest.Server {
	var db database2.DatabaseInterface
	var configuration models.Configuration
	var dataModel []initialisation.DataModel
	api := &core.Api{Json: initialisation.JsonHandler{File: "../../config/config.test.json"}}
	api.Initialisation(&configuration, &dataModel)
	db = &database2.MongoDB{Name: configuration.Db.Name, Url: configuration.Db.Url}
	router := mux.NewRouter()
	router.Use(logging.Logging())
	controller.InitControllers(router, &configuration, &dataModel, db)
	server := httptest.NewServer(router)
	return server
}

func TestServerIntegrationHealth(t *testing.T) {
	server := startServer()
	defer server.Close()
	resp, err := http.Get(server.URL + "/health")
	if err != nil {
		t.Fatalf("failed to make GET request to server: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d want %d", resp.StatusCode, http.StatusOK)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respBody := buf.String()
	expectedRespBody := "This is the health endpoint\n"
	if respBody != expectedRespBody {
		t.Errorf("unexpected response body: got %q want %q", respBody, expectedRespBody)
	}
}

func TestServerIntegrationBook(t *testing.T) {
	server := startServer()
	defer server.Close()
	resp, err := http.Get(server.URL + "/book/read")
	if err != nil {
		t.Fatalf("failed to make GET request to server: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("unexpected status code: got %d want %d", resp.StatusCode, http.StatusOK)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respBody := buf.String()
	expectedRespBody := "404 page not found\n"
	if respBody != expectedRespBody {
		t.Errorf("unexpected response body: got %q want %q", respBody, expectedRespBody)
	}
}
