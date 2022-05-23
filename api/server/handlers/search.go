package handlers

import (
	"github.com/georgejdanforth/crate-digger/database"
	"github.com/georgejdanforth/crate-digger/models"
	"github.com/georgejdanforth/crate-digger/server/queries"
	"github.com/gin-gonic/gin"
	"log"
)

const SimilarityThreshold = 0.2

func Search(ctx *gin.Context) {
	var search models.Search
	if ctx.ShouldBindQuery(&search) != nil {
		// TODO: respond with an actual error
		ctx.Data(400, gin.MIMEJSON, nil)
	}
	switch(search.Entity) {
	case "artist":
		{
			results := []models.ArtistSearchResult{}
			doSearch(ctx, &results, queries.SearchArtist, search.Query)
		}
	case "label":
		{
			results := []models.LabelSearchResult{}
			doSearch(ctx, &results, queries.SearchLabel, search.Query)
		}
	default:
		ctx.Data(400, gin.MIMEJSON, nil)
	}
}

func doSearch[T models.SearchResult](ctx *gin.Context, results *[]T, sql string, query string) {
	db := database.GetDb()
	if err := db.Select(results, sql, query, SimilarityThreshold); err == nil {
		ctx.JSON(200, models.SearchResults[T]{Results: *results})
	} else {
		// TODO: proper error logging and return
		log.Printf("ERROR: %v", err)
		ctx.Data(500, gin.MIMEJSON, nil)
	}
}
