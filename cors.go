package VAS_util_go

import "net/http"

func enableCors(w *http.ResponseWriter, corsUrl string) {
	(*w).Header().Set("Access-Control-Allow-Origin", corsUrl)
}
