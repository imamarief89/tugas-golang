package controllers

import (
	"GOLANG-IMAM/entities"
	"GOLANG-IMAM/models"
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

var taksModel = models.New()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/index.html")
	temp.Execute(w, data)
}

func GetData() string {

	buffer := &bytes.Buffer{}

	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/data.html")

	var taks []entities.Taks
	err := taksModel.FindAll(&taks)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"taks": taks,
	}

	temp.ExecuteTemplate(buffer, "data.html", data)

	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}
	var taks entities.Taks

	if err != nil {
		data = map[string]interface{}{
			"title": "Tambah Data Taks",
			"taks":  taks,
		}
	} else {

		err := taksModel.Find(id, &taks)
		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"title": "Edit Data Taks",
			"taks":  taks,
		}
	}

	temp, _ := template.ParseFiles("views/form.html")
	temp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		r.ParseForm()
		var taks entities.Taks

		taks.TaksName = r.Form.Get("taks_name")
		taks.Assignor = r.Form.Get("assignor")
		taks.Deadline = r.Form.Get("deadline")
		taks.Finish = r.Form.Get("finish")

		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		var data map[string]interface{}

		if err != nil {
			// insert data
			err := taksModel.Create(&taks)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data berhasil disimpan",
				"data":    template.HTML(GetData()),
			}
		} else {
			// mengupdate data
			taks.Id = id
			err := taksModel.Update(taks)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data berhasil diubah",
				"data":    template.HTML(GetData()),
			}
		}

		ResponseJson(w, http.StatusOK, data)
	}
}

func Done(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	err = taksModel.Done(id)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"message": "Data is Done",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	err = taksModel.Delete(id)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"message": "Data berhasil dihapus",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"error": message})
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
