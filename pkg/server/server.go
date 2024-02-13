package server

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/OkSProject/OkServer/pkg/config"
)

func Run() {
	logDir := config.GetLogDir()
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating log directory %s: %e", logDir, err)
	}

	logFile, err := os.OpenFile(logDir+"oksi-http.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	defer logFile.Close()

	// Allows for both writing a log file and printing it to the command line.
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	// Serves site(s) from directory "/".
	fileServer := http.FileServer(http.Dir(config.GetAssetDir()))
	mod_php := modules.mod_php()
	mod_py := modules.mod_py()
	mod_go := modules.mod_go()
	http.Handle("/", fileServer)
	http.HandleFunc("/php/", mod_php)
	http.HandleFunc("/py/", mod_py)
	http.HandleFunc("/go/", mod_go)

	port := config.GetListenPort()
	log.Printf("OkS! Running on :%s...\n", port)

	// Error checking
	// Sets "err" to the output of the HTTP server when it starts.
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Error: ", err)
	}
}
