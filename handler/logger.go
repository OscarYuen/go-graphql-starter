package handler

import (
	"bytes"
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http"
	"os"
)

type LoggerHandler struct {
	DebugMode bool
	Logger    *logging.Logger
}

func NewLogger(appName *string, debugMode bool, logFormat *string) *logging.Logger {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter(*logFormat)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.INFO, "")
	if debugMode {
		backendLeveled.SetLevel(logging.DEBUG, "")
	}

	logging.SetBackend(backendLeveled)
	logger := logging.MustGetLogger(*appName)
	return logger
}

func (l *LoggerHandler) Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Logger.Infof("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto)
		l.Logger.Infof("User agent : %s", r.UserAgent())
		if l.DebugMode {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				l.Logger.Errorf("Reading request body error: %s", err)
			}
			reqStr := ioutil.NopCloser(bytes.NewBuffer(body))
			l.Logger.Debugf("Request body : %v", reqStr)
			r.Body = reqStr
		}
		h.ServeHTTP(w, r)
	})
}
