package server

import (
	"example.com/hello/pkg/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Run() {

	// Creates /log directory to store server logs.
	// Exits if an error occurs with the directory creation.
	logDir := config.GetLogDir()
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating log directory %s: %e", logDir, err)
	}

	// Creates or opens `server.log` for writing purposes.
	// Exits if an error occurs while opening the log file.
	// Closes the log file once done and sets log_file as the output.
	logFile, err := os.OpenFile(logDir+"okserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	defer logFile.Close()

	// Allows for both writing a log file and printing it to the command line.
	mwLog := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(mwLog)

	// Serves site(s) from directory "/".
	dir := config.GetAssetDir()
	fileServer := http.FileServer(http.Dir(dir))
	http.Handle("/", fileServer)

	// Sets server to listen on Port 8080.
	port := config.GetListenPort()

	// Starts server.
	log.Printf("OkServer! Running on: %s...\n", port)

	// Error checking
	// Sets "err" to the output of the HTTP server when it starts.
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Println("Error: ", err)
	}
}