package main

import (
	"fmt"
	"net/http"
	"log"
	"os/exec"
	"time"
	"os"
	"strings"
	"errors"
)

func gitPull() {
	log.Println("Git pull triggered, running...")

	out, err := exec.Command("git", "pull").CombinedOutput()
	if (err != nil) {
		log.Println("Error calling git pull: " + err.Error())
		log.Print("\n" + string(out))
	}
}

func getEnv(name string) (string, error) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == name {
			return pair[1], nil
		}
	}

	return "", errors.New("no such environment variable")
}

func pullEndpoint(w http.ResponseWriter, r *http.Request) {
	// parse query string
	r.ParseForm()

	unauthorizedMsg := "You are not authorized to perform this action"

	// extract auth parameter
	auth := r.Form["auth"]
	if len(auth) < 1 {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, unauthorizedMsg)
		return
	}

	// verify auth parameter
	validKeys, err := getEnv("GIT_PULL_AUTH")
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal server error")
		fmt.Println("Error fetching environment variable: " + err.Error())
		return
	}

	for _, key := range strings.Split(validKeys, ";") {
		if key == auth[0] {
			go gitPull()

			w.WriteHeader(http.StatusAccepted);
			fmt.Fprint(w, "Git pull triggered!");
			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, unauthorizedMsg)
}

func main() {
	log.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " [AAA] ")

	http.HandleFunc("/pull", pullEndpoint)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
