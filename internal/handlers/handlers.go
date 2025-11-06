package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler возвращает HTML из файла index.html для корневого эндпоинта /
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	file, err := os.Open("index.html")
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, "file index.html not found: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = io.Copy(w, file)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, "error reading the file: "+err.Error(), http.StatusInternalServerError)
	}
}

// UploadHandler обрабатывает загрузку файла на эндпоинте /upload
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "error when parsing the form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "file was not found in the form: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	inputText := string(data)

	result, err := service.Convert(inputText)
	if err != nil {
		http.Error(w, "conversion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Генерируем имя файла на основе текущего времени
	timestamp := time.Now().UTC()
	formatted := timestamp.Format("2006-01-02_15-04-05.999999")
	filename := formatted + ".txt"

	// Создаём и записываем результат в локальный файл
	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "error when creating the file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.WriteString(result)
	if err != nil {
		http.Error(w, "error when writing to a file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"result": "` + result + `"}`))
}
