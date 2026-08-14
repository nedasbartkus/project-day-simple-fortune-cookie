package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	fortuneH := &fortuneHandler{store: &datastoreDefault}
	mux.Handle("/fortunes", fortuneH)
	mux.Handle("/fortunes/", fortuneH)
	return mux
}

func TestRootHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	setupMux().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
	if rr.Body.String() != "ok" {
		t.Fatalf("expected body %q, got %q", "ok", rr.Body.String())
	}
}

func TestUnknownPathStillNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rr := httptest.NewRecorder()

	setupMux().ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rr.Code)
	}
}
