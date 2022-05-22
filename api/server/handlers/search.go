package handlers

import (
	"github.com/georgejdanforth/crate-digger/database"
	"github.com/georgejdanforth/crate-digger/models"
	"github.com/georgejdanforth/crate-digger/server/queries"
	"github.com/gin-gonic/gin"
	"log"
)

func Search(ctx *gin.Context) {
	db := database.GetDb()
	var search models.Search
	if ctx.ShouldBindQuery(&search) != nil {
		// TODO: respond with an actual error
		ctx.Data(400, gin.MIMEJSON, nil)
	}
	switch(search.Entity) {
	case "label":
		{
			results := []models.LabelSearchResult{}
			
			if err := db.Select(&results, queries.SearchLabel, search.Query, 0.2); err == nil {
				ctx.JSON(200, models.LabelSearchResults{Results: results})
			} else {
				log.Printf("ERROR: %v", err)
				ctx.Data(500, gin.MIMEJSON, nil)
			}
		}
	default:
		ctx.Data(400, gin.MIMEJSON, nil)
	}
}
