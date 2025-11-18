package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl *template.Template

func init() {
	// Инициализируем шаблон при запуске
	var err error
	tmpl, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatal("Ошибка загрузки шаблонов: ", err)
	}
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "homepage.html", nil)
	} else {
		data := map[string]interface{}{
			"Error":  "",
			"Result": "",
		}
		r.ParseForm()
		num1Str := r.FormValue("num1")
		num2Str := r.FormValue("num2")
		operation := r.FormValue("operation")

		num1, err := strconv.ParseFloat(num1Str, 64)
		if err != nil {
			data["Error"] = "incorrect number"
			tmpl.ExecuteTemplate(w, "homepage.html", data)
		}

		num2, err := strconv.ParseFloat(num2Str, 64)
		if err != nil {
			data["Error"] = "incorrect number"
			tmpl.ExecuteTemplate(w, "homepage.html", data)
		}

		var result float64
		var errorMsg string

		switch operation {
		case "add":
			result = num1 + num2
		case "subtract":
			result = num1 - num2
		case "multiply":
			result = num1 * num2
		case "divide":
			if num2 == 0 {
				data["Error"] = "Деление на ноль невозможно"
				tmpl.ExecuteTemplate(w, "homepage.html", data)
				return
			}
			result = num1 / num2
		default:
			errorMsg = "Неверная операция"
			data["Error"] = errorMsg
		}

		if errorMsg == "" {
			data["Result"] = strconv.FormatFloat(result, 'f', -1, 64)
		}
		tmpl.ExecuteTemplate(w, "homepage.html", data)
	}
}
