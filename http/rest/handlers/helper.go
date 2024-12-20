package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/perdanaph/todoApiGo/pkg/erru"
)

func (s service) respond(w http.ResponseWriter, data interface{}, status int) {
	var respData interface{}
	switch v := data.(type) {
	case nil:
	case erru.ErrArgument:
		status = http.StatusBadRequest
		respData = ErrorResponse{ErrorMessage: v.Unwrap().Error()}
	case error:
		if http.StatusText(status) == "" {
			status = http.StatusInternalServerError
		} else {
			respData = ErrorResponse{ErrorMessage: v.Error()}
		}
	default:
		respData = data
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		err := json.NewEncoder(w).Encode(respData)
		if err != nil {
			http.Error(w, "Could not encode in json", http.StatusBadRequest)
			return
		}
	}
}

func (s service) decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// it reads to the memory.
func (s service) readRequestBody(r *http.Request) ([]byte, error) {
	// Read the content
	var bodyBytes []byte
	var err error
	if r.Body != nil {
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			err := errors.New("could not read request body")
			return nil, err
		}
	}
	return bodyBytes, nil
}

func (s service) restoreRequestBody(r *http.Request, bodyBytes []byte) {

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
}
