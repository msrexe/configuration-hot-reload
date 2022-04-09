package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var logger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: &logrus.JSONFormatter{},
	Level:     logrus.InfoLevel,
}

var logConfig = &LogConfig{}

func main() {
	server := http.NewServeMux()

	loadConfigFromJSON()

	server.HandleFunc("/test", testHandler)
	server.HandleFunc("/log-config", logConfigHandler)

	logger.Info("Starting server")
	logger.Fatal(http.ListenAndServe(":8080", server))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	logger.Debug(fmt.Sprintf("%+v", r))
	w.WriteHeader(http.StatusOK)
}

func logConfigHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.Unmarshal(body, logConfig)
		logConfig.Update(*logConfig)

		logger.Info("Log config updated")
		logger.Debug(fmt.Sprintf("%+v", logConfig))
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.Write(logConfig.GetConfigAsJSON())
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.Error(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	}
}

func loadConfigFromJSON() error {
	rawConfig, _ := ioutil.ReadFile("config.json")

	return json.Unmarshal(rawConfig, &logConfig)
}
