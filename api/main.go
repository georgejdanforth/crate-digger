package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"encoding/json"
	"log"
)

type NullInt16 sql.NullInt16

func (ni *NullInt16) Scan(value interface{}) error {
	return (*sql.NullInt16)(ni).Scan(value)
}

func (ni *NullInt16) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int16)
}

type NullInt32 sql.NullInt32

func (ni *NullInt32) Scan(value interface{}) error {
	return (*sql.NullInt32)(ni).Scan(value)
}

func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int32)
}

const SearchLabel = `
SELECT
	label.id               AS id,
	label.gid              AS gid,
	label.name             AS name,
	label.begin_date_year  AS begin_date_year,
	label.begin_date_month AS begin_date_month,
	label.begin_date_day   AS begin_date_day,
	label.end_date_year    AS end_date_year,
	label.end_date_month   AS end_date_month,
	label.end_date_day     AS end_date_day,
	label.label_code       AS label_code,
	label.type             AS type,
	label.comment          AS comment,
	MAX(name_scores.score) AS score
FROM (
	SELECT name, similarity(name, $1) AS score
	FROM (
		SELECT name              FROM label       UNION ALL
		SELECT name              FROM label_alias UNION ALL
		SELECT sort_name AS name FROM label_alias
	) names
	WHERE similarity(name, $1) > $2 
	ORDER BY score DESC
	LIMIT 100
) name_scores
LEFT JOIN label_alias AS alias ON (alias.name = name_scores.name OR alias.sort_name = name_scores.name)
INNER JOIN label ON (alias.label = label.id OR name_scores.name = label.name)
GROUP BY label.id, label.gid, label.name
ORDER BY score DESC, label.name
LIMIT 50;`

type LabelSearchResult struct {
	Id             int       `db:"id" json:"id"`
	Gid            string    `db:"gid" json:"gid"`
	Name           string    `db:"name" json:"name"`
	BeginDateYear  NullInt16 `db:"begin_date_year" json:"beginDateYear"`
	BeginDateMonth NullInt16 `db:"begin_date_month" json:"beginDateMonth"`
	BeginDateDay   NullInt16 `db:"begin_date_day" json:"beginDateDay"`
	EndDateYear    NullInt16 `db:"end_date_year" json:"endDateYear"`
	EndDateMonth   NullInt16 `db:"end_date_month" json:"endDateMonth"`
	EndDateDay     NullInt16 `db:"end_date_day" json:"endDateDay"`
	LabelCode      NullInt32 `db:"label_code" json:"labelCode"`
	Type           NullInt32 `db:"type" json:"type"`
	Comment        string    `db:"comment" json:"comment"`
	Score          float32   `db:"score" json:"score"`
}

type LabelSearchResults struct {
	Results []LabelSearchResult `json:"results"`
}

type Search struct {
	Entity string `form:"entity"`
	Query  string `form:"query"`
}

func main() {
	db := sqlx.MustConnect("postgres", "host=localhost port=15432 user=musicbrainz password=musicbrainz dbname=musicbrainz sslmode=disable")
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/healthz", func(c *gin.Context) {
			c.Data(200, gin.MIMEJSON, nil)
		})

		v1 := api.Group("/v1")
		{
			v1.GET("/search", func(c *gin.Context) {
				var search Search
				if c.ShouldBindQuery(&search) != nil {
					// TODO: respond with an actual error
					c.Data(400, gin.MIMEJSON, nil)
				}

				switch(search.Entity) {
				case "label":
					{
						results := []LabelSearchResult{}
						
						if err := db.Select(&results, SearchLabel, search.Query, 0.2); err == nil {
							c.JSON(200, LabelSearchResults{Results: results})
						} else {
							log.Printf("ERROR: %v", err)
							c.Data(500, gin.MIMEJSON, nil)
						}
					}
				default:
					c.Data(400, gin.MIMEJSON, nil)
				}
			})
		}
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
