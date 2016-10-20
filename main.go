package main

import (
	"net/http"
	"fmt"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3005"
	}

	host := os.Getenv("REDIRECT_TO")
	if host == "" {
		fmt.Fprintf(os.Stderr, "REDIRECT_TO environment variable not set;\n")
		fmt.Fprintf(os.Stderr, "do you need to `cf set-env REDIRECT_TO host-or-ip'?\n")
		os.Exit(1)
	}
	scheme := os.Getenv("REDIRECT_SCHEME")
	if scheme == "" {
		scheme = "https"
	}
	if scheme != "http" && scheme != "https" {
		fmt.Fprintf(os.Stderr, "REDIRECT_SCHEME set to invalid value of '%s'\n")
		os.Exit(1)
	}

	status_s := os.Getenv("STATUS_3XX")
	if status_s == "" {
		status_s = "302"
	}
	status, err := strconv.Atoi(status_s)
	if err != nil || status < 300 || status > 399 {
		fmt.Fprintf(os.Stderr, "STATUS_3XX set to a non-3xx value of '%s'\n", status_s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "(more specific error was: %s)\n", err)
		}
		os.Exit(1)
	}

	debug := os.Getenv("DEBUG") != ""

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if debug {
			fmt.Fprintf(os.Stderr, "redirecting %s\n", r.URL.String())
		}
		r.URL.Scheme = scheme
		r.URL.Host = host
		w.Header().Add("Location", r.URL.String())
		if debug {
			fmt.Fprintf(os.Stderr, "   (%d) -> %s\n", status, r.URL.String())
		}
		w.WriteHeader(status)
	})
	if debug {
		fmt.Printf("starting up on *:%s\n", port)
		fmt.Printf("redirecting all requests to %s:%s\n", scheme, host)
	}
	err = http.ListenAndServe(":"+port, nil)
	fmt.Fprintf(os.Stderr, "server aborted: %s\n", err)
	fmt.Fprintf(os.Stderr, "shutting down.\n")
	os.Exit(2)
}
