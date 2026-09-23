package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	root, err := os.OpenRoot("static")

	if err != nil {
		log.Printf("error opening root: %v", err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	defer root.Close()

	htmlFile, err := root.Open("index.html")

	if err != nil {
		log.Printf("html file not found: %v", err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	defer htmlFile.Close()

	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, htmlFile)

	if err != nil {
		log.Printf("send error: %v", err)
		http.Error(w, "ошибка при отправке файла", http.StatusInternalServerError)
		return
	}
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")

	if err != nil {
		log.Printf("get file error: %v", err)
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	result := service.ParseData(string(data))

	fmt.Println(result)

	root, err := os.OpenRoot("uploads")
	if err != nil {
		log.Printf("open upload directory error: %v", err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	fileName := time.Now().UTC().String()
	extension := filepath.Ext(handler.Filename)

	dst, err := root.Create(fmt.Sprintf("%s%s", fileName, extension))
	if err != nil {
		log.Printf("create file error: %v", err)
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.WriteString(dst, result)
	if err != nil {
		log.Printf("copy file error: %v", err)
		http.Error(w, "ошибка при записи файла", http.StatusInternalServerError)
		return
	}
}
