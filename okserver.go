/*
	 OkServer! (OkS!) - An ok, lightweight, clean, and straightforward HTTP server!
    Copyright (C) 2024  Cole Rathbun <lowcr@proton.me>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// Creates /log directory to store server logs.
	// Exits if an error occurs with the directory creation.
	err := os.MkdirAll("./log", os.ModePerm)
	if err != nil {
		log.Fatal("Error creating /log directory: ", err)
	}

	// Creates or opens `server.log` for writing purposes.
	// Exits if an error occurs while opening the log file.
	// Closes the log file once done and sets log_file as the output.
	log_file, err := os.OpenFile("./log/okserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	defer log_file.Close()

	// Allows for both writing a log file and printing it to the command line.
	mw_log := io.MultiWriter(log_file, os.Stdout)
	log.SetOutput(mw_log)

	// Serves site(s) from directory "/".
	dir := "./www/html/"
	file_server := http.FileServer(http.Dir(dir))
	http.Handle("/", file_server)

	// Sets server to listen on Port 8080.
	port := 8080

	// Starts server.
	log.Printf("OkServer! Running on: %d...\n", port)

	// Error checking
	// Sets "err" to the output of the HTTP server when it starts.
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Println("Error: ", err)
	}
}
