package integrationTests

import (
	"bytes"
	"github.com/gorilla/mux"
	"httpServer/src/controller"
	"httpServer/src/core"
	"httpServer/src/initialisation"
	"httpServer/src/middlewares/logging"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerIntegration(t *testing.T) {
	_ = &core.Api{Json: initialisation.JsonHandler{File: "config/config.json"}}
	router := mux.NewRouter()
	router.Use(logging.Logging())
	controller.InitControllers(router)
	server := httptest.NewServer(router)
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
