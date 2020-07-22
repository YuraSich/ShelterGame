package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting Server...")

	router := mux.NewRouter()

	router.HandleFunc("/", handler)

	log.Info("Server Listen And Serve localhost:8080")
	http.ListenAndServe(":8080", router)
}

var cookies = map[string]*securecookie.SecureCookie{
	"previous": securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	),
	"current": securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	),
}

func handler(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"foo": "bar",
	}
	if encoded, err := securecookie.EncodeMulti("cookie-name", value, cookies["current"]); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie-name",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
	ReadCookieHandler(w, r)
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("cookie-name"); err == nil {
		value := make(map[string]string)
		err = securecookie.DecodeMulti("cookie-name", cookie.Value, &value, cookies["current"], cookies["previous"])
		if err == nil {
			fmt.Fprintf(w, "The value of foo is %q", value["foo"])
		}
	}
}
