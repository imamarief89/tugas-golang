package main

import (
	"GOLANG-IMAM/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/get_form", controllers.GetForm)
	http.HandleFunc("/store", controllers.Store)
	http.HandleFunc("/done", controllers.Done)
	http.HandleFunc("/delete", controllers.Delete)
	http.ListenAndServe(":8080", nil)
}
