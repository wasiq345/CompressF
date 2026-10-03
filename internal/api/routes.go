package api

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/health", Health)
	//mux.HandleFunc("POST /api/uploadfile", UploadFile)
	mux.HandleFunc("POST /api/compressfile", CompressFile)
	//mux.HandleFunc("POST /api/convert", ConvertFile)
}
