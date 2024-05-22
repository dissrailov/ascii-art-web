package handlers

import (
	"ascii-art-web/ascii"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type errors struct {
	ErrorCode int
	ErrorMsg  string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func GeneratePage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if r.URL.Path != "/generate" {
		w.WriteHeader(http.StatusNotFound)
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	input := r.FormValue("textInput")
	input = strings.ReplaceAll(input, "\r\n", "\\n")

	if len(input) < 1 || len(input) > 150 {
		w.WriteHeader(http.StatusBadRequest)
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if ascii.AsciiCheck(input) != nil {
		w.WriteHeader(http.StatusBadRequest)
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	filename := r.FormValue("filename")
	// fmt.Println(filename)
	if filename != "standard" && filename != "shadow" && filename != "thinkertoy" {

		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	// end....
	output, error := ascii.AsciiMain(input, filename)
	if error != nil {
		// use error handler
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	ts.Execute(w, output)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, errCode int, msg string) {
	t, err := template.ParseFiles("templates/Error.html")
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	Errors := errors{
		ErrorCode: errCode,
		ErrorMsg:  msg,
	}
	// w.WriteHeader(Errors.ErrorCode)
	t.Execute(w, Errors)
}
