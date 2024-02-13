package modules

import (
	"net/http"
	"os/exec"
	"strings"
)

func mod_php(w http.ResponseWriter, r *http.Request) {
	// Sets up PHP path by taking the URL request and trimming out everything except for the file.
	php_path := strings.TrimPrefix(r.URL.Path, "/php/")

	// Sets up and executes the PHP command on the server's machine.
	// Returns an Internal Server Error if PHP isn't installed on the user's machine.
	run_php := exec.Command("php", php_path)
	output, err := run_php.Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Writes the output of the request to the user.
	w.Header().Set("Content-Type", "text/html")
	w.Write(output)
}

func mod_py(w http.ResponseWriter, r *http.Request) {
	// Sets up Python path by taking the URL request and trimming out everything except for the file.
	py_path := strings.TrimPrefix(r.URL.Path, "/py/")

	// Sets up and executes the PHP command on the server's machine.
	// Returns an Internal Server Error if PHP isn't installed on the user's machine.
	run_py := exec.Command("py", py_path)
	output, err := run_py.Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Writes the output of the request to the user.
	w.Header().Set("Content-Type", "text/html")
	w.Write(output)
}

func mod_go(w http.ResponseWriter, r *http.Request) {
	go_path := strings.TrimPrefix(r.URL.Path, "/go/")
	http.ServeFile(w, r, go_path)
}
