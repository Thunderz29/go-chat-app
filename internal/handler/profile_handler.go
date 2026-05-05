package handler

import "net/http"

func Profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("profile endpoint (protected)"))
}
