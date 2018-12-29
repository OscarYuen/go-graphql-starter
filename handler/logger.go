package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/op/go-logging"
)

type LoggerHandler struct {
	DebugMode bool
}

func (l *LoggerHandler) Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := ctx.Value("log").(*logging.Logger)
		log.Infof("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto)
		log.Infof("User agent : %s", r.UserAgent())
		if l.DebugMode {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Errorf("Reading request body error: %s", err)
			}
			reqStr := ioutil.NopCloser(bytes.NewBuffer(body))
			log.Debugf("Request body : %v", reqStr)
			r.Body = reqStr
		}
		h.ServeHTTP(w, r)
	})
}
