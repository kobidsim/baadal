package handler

import (
	"fmt"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "error uploading file", http.StatusInternalServerError)
	}
	if r.MultipartForm == nil {
		http.Error(w, "no files uploaded", http.StatusBadRequest)
		return
	}
	files := r.MultipartForm.File
	if len(files) == 0 {
		http.Error(w, "no files uploaded", http.StatusBadRequest)
	}

	for _, file := range files["files"] {
		if file != nil {
			fmt.Fprintf(w, "filename: %v", file.Filename)
		}
	}
}
