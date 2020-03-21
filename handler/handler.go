package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/MrJootta/GoUrl/internal/utils"

	"github.com/MrJootta/GoUrl/internal/storage"
)

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"response"`
}

type handler struct {
	storage storage.Service
}

// New returns handler
func New(storage storage.Service) http.Handler {
	h := handler{storage}

	mux := http.NewServeMux()

	mux.HandleFunc("/encode/", responseHandler(h.encode))
	mux.HandleFunc("/info/", responseHandler(h.info))
	mux.HandleFunc("/", h.redirect)

	return mux
}

func responseHandler(h func(io.Writer, *http.Request) (interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)
		if err != nil {
			data = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		err = json.NewEncoder(w).Encode(response{Data: data, Status: status})
		if err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}

func (h handler) encode(w io.Writer, r *http.Request) (interface{}, int, error) {
	if r.Method != http.MethodPost {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("method %s not allowed", r.Method)
	}

	var input struct{ URL string }
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode request: %v", err)
	}

	url := strings.TrimSpace(input.URL)
	if url == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("URL is empty")
	}

	if !strings.Contains(url, "http") {
		url = "http://" + url
	}

	code, err := utils.GenerateCode()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("did not generated code: %s", err.Error())
	}

	c, err := h.storage.NewCode(code, url)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("did not store in database: %v", err)
	}

	return c, http.StatusCreated, nil
}

func (h handler) info(w io.Writer, r *http.Request) (interface{}, int, error) {
	if r.Method != http.MethodGet {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("Method %s not allowed", r.Method)
	}

	code := r.URL.Path[len("/info/"):]

	model, err := h.storage.CodeInfo(code)
	if err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("URL not found")
	}

	return model, http.StatusOK, nil
}

func (h handler) redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	code := r.URL.Path[len("/"):]

	url, err := h.storage.GetURL(code)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("URL Not Found"))
		return
	}

	http.Redirect(w, r, url.URL, http.StatusMovedPermanently)
}
