package routes

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bienvenido a la API de Tareas v1"))
}
