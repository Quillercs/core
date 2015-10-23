package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
	Platform string `json:"platform"`
	Version  string `json:"version"`
}

func CheckEngine(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		req := Engine{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		switch req.Platform {
		case "freeswitch":
			switch req.Version {
			case "1.6.2":
				log.Printf("FreeSwitch Version: %s supported", req.Version)

			default:
				err := fmt.Sprintf("FreeSwitch Version: %s not supported yet :(", req.Version)
				log.Println(err)
				http.Error(w, err, 400)
				return

			}
		case "asterisk":
			err := fmt.Sprintln("Asterisk is not supported yet")
			log.Println(err)
			http.Error(w, err, 400)
			return

		default:
			err := fmt.Sprintf("Engine %s is not supported", req.Platform)
			log.Println(err)
			http.Error(w, err, 400)
			return
		}

		next.ServeHTTP(w, r)
	})
}
