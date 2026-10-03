package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var ErrFromFile = errors.New("Unable to get File")

func CreateFile(r *http.Request) (*os.File, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFromFile, err)
	}
	defer file.Close()
	temp, err := os.CreateTemp("", "userfile-*")
	if err != nil {
		return nil, fmt.Errorf("create uploaded file %w", err)
	}
	_, err = io.Copy(temp, file)
	if err != nil {
		defer os.Remove(temp.Name())
		defer temp.Close()
		return nil, fmt.Errorf("copy uploaded file %w", err)
	}
	return temp, nil
}

func CompressFile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	temp, err := CreateFile(r)
	defer os.Remove(temp.Name())
	defer temp.Close()
	if err != nil {
		if errors.Is(err, ErrFromFile) {
			ResponseWithErr(w, http.StatusBadRequest, err.Error())
		} else {
			ResponseWithErr(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

}
