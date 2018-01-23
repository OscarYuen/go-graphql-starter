package handler

import (
	"bytes"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"os"
)

type Logger struct {
	AppName   *string
	DebugMode bool
	LogFormat *string
}

func (l *Logger) Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backend := logging.NewLogBackend(os.Stderr, "", 0)
		format := logging.MustStringFormatter(*l.LogFormat)
		backendFormatter := logging.NewBackendFormatter(backend, format)

		backendLeveled := logging.AddModuleLevel(backendFormatter)
		backendLeveled.SetLevel(logging.INFO, "")
		if l.DebugMode {
			backendLeveled.SetLevel(logging.DEBUG, "")
		}

		logging.SetBackend(backendLeveled)
		logger := logging.MustGetLogger(*l.AppName)
		logger.Infof("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto)
		logger.Infof("User agent : %s", r.UserAgent())
		if l.DebugMode {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logger.Errorf("Reading request body error: %s", err)
			}
			reqStr := ioutil.NopCloser(bytes.NewBuffer(body))
			logger.Debugf("Request body : %v", reqStr)
			r.Body = reqStr
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "logger", logger)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
