package api

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi(t *testing.T) {
	handler := MazeApiRouter()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/").Expect().Status(http.StatusOK).Body().NotEmpty()
	e.GET("/0/0/0").Expect().Status(http.StatusOK).Body().NotEmpty()
	e.GET("/DepthFirstGenerator/RecursiveSolver").Expect().Status(http.StatusOK).Body().NotEmpty()
	e.GET("/BreadthFirstGenerator/ConcurrentSolver").Expect().Status(http.StatusOK).Body().NotEmpty()

	e.GET("/xyz/RecursiveSolver").Expect().Status(http.StatusBadRequest).Body().Contains("xyz").Contains("generator")
	e.GET("/DepthFirstGenerator/abc").Expect().Status(http.StatusBadRequest).Body().Contains("abc").Contains("solver")

	e.GET("/abc").Expect().Status(http.StatusNotFound)
}
