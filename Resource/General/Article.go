package General

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func Article(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "POST" {
			sQuery := r.FormValue("sQuery")

			result := bson.M{}

			_, err := db.Exec(sQuery)
			if err != nil {
				result["status"] = 0
				result["note"] = err.Error()
				result["items"] = []bson.M{}
			} else {
				result["status"] = 1
				result["note"] = ""
				result["items"] = []bson.M{}
			}
			json.NewEncoder(w).Encode(result)
		}

		if r.Method == "GET" {
			sQuery := r.FormValue("sQuery")

			result := bson.M{}

			_, err := db.Exec(sQuery)
			if err != nil {
				result["status"] = 0
				result["note"] = err.Error()
				result["items"] = []bson.M{}
			} else {
				result["status"] = 1
				result["note"] = ""
				result["items"] = []bson.M{}
			}
			json.NewEncoder(w).Encode(result)
		}
	})
}
