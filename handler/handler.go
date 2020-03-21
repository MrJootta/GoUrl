package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/MrJootta/GoUrl/internal/storage"
	"github.com/MrJootta/GoUrl/internal/utils"
)

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"response"`
}

type requestPost struct {
	URL string `json:"url"`
}

type handler struct {
	storage storage.Service
}

// New returns handler
func New(storage storage.Service) http.Handler {
	h := handler{storage}

	mux := http.NewServeMux()

	mux.HandleFunc("/create", wrapper(h.create))
	mux.HandleFunc("/info", wrapper(h.info))
	mux.HandleFunc("/", h.redirect)

	return mux
}

func wrapper(h func(io.Writer, *http.Request) (interface{}, int, error)) http.HandlerFunc {
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

func (h handler) create(w io.Writer, r *http.Request) (interface{}, int, error) {
	if r.Method != http.MethodPost {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("%s not allowed", r.Method)
	}

	var req requestPost

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode: %v", err)
	}

	urlParam, err := url.Parse(req.URL)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("URL not valid")
	}

	code, err := utils.GenerateCode()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("did not generated code: %s", err.Error())
	}

	_, err = h.storage.NewCode(code, urlParam.String())
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("did not store in database: %v", err)
	}

	returnData := struct {
		Url  string `json:"url"`
		Code string `json:"code"`
	}{
		utils.URLBuilder(code),
		code,
	}

	return returnData, http.StatusCreated, nil
}

func (h handler) info(w io.Writer, r *http.Request) (interface{}, int, error) {
	if r.Method != http.MethodGet {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("%s not allowed", r.Method)
	}

	code := r.URL.Path[len("/info/"):]

	model, err := h.storage.CodeInfo(code)
	if err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("code not found, try to encode new url")
	}

	returnData := struct {
		Url            string `json:"url"`
		Code           string `json:"code"`
		NumberOfVisits int    `json:"number_of_visits"`
	}{
		utils.URLBuilder(code),
		code,
		len(model),
	}

	return returnData, http.StatusOK, nil
}

func (h handler) redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("endpoint only available over GET method"))

		return
	}

	code := r.URL.Path[len("/"):]

	urlParam, err := h.storage.GetURL(code)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("code not found, try to encode new url"))

		return
	}

	http.Redirect(w, r, urlParam.URL, http.StatusMovedPermanently)
}
