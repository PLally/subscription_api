package httphelpers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StatusResponse(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	w.Write([]byte(http.StatusText(status)))

}

func Response(w http.ResponseWriter, status int, text string) {
	w.WriteHeader(status)
	w.Write([]byte(text))
}

func WriteJson(w http.ResponseWriter, obj interface{}, status int) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	w.WriteHeader(status)
	w.Write(data)
}