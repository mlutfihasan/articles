package resource

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/mlutfihasan/articles/config"
)

type Article struct {
	Id      int       `json:"id"`
	Author  string    `json:"author"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

func Articles(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var articleData Article
			err = json.Unmarshal(body, &articleData)
			if err != nil {
				panic(err)
			}

			query := `INSERT INTO articles(id, author, title, body, created)
						SELECT ` + strconv.Itoa(articleData.Id) + `, '` + articleData.Author + `', '` + articleData.Title + `', '` + articleData.Body + `', NOW();`
			_, err = db.Exec(query)
			var result config.UpdateResult
			if err != nil {
				result = config.UpdateResult{
					Status: "0",
					Note:   err.Error(),
				}
			} else {
				result = config.UpdateResult{
					Status: "1",
					Note:   "",
				}
			}

			json.NewEncoder(w).Encode(result)
			return
		}

		if r.Method == "GET" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var articleData Article
			err = json.Unmarshal(body, &articleData)
			if err != nil {
				panic(err)
			}

			query := `SELECT * FROM articles
						WHERE title LIKE '%` + articleData.Title + `%'
						AND body LIKE '%` + articleData.Body + `%'
						AND author LIKE '%` + articleData.Author + `%'
						ORDER BY created DESC;`

			result := config.SQLToJSON(db, query)

			json.NewEncoder(w).Encode(result)
			return
		}
	})
}
