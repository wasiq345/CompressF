package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) error {
	resp, err := json.Marshal(payload)
	if err != nil {
		return errors.New("Unable to Marshal Data")
	}
	w.WriteHeader(code)
	w.Write(resp)
	return nil
}

func ResponseWithErr(w http.ResponseWriter, code int, text string) error {
	err := ResponseWithJson(w, code, map[string]string{"error: ": text})
	if err != nil {
		return err
	}
	return nil
}
