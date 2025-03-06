package controllers

import (
	"mime/multipart"
	"net/http"
)

type Gallery struct {
	File multipart.File
	FileHeader *multipart.FileHeader
	GalleryPath string
	Templates struct {
		New Template
	}
}

func (g Gallery) New(w http.ResponseWriter, r *http.Request) {
	g.Templates.New.Execute(w, g)
}

func (g Gallery) Upload(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}
	g.File = file
	g.FileHeader = fileHeader
	g.GalleryPath = "./images/" + fileHeader.Filename
	w.Write([]byte("File uploaded successfully"))
}
