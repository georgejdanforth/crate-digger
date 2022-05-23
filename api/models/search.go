package models

type Search struct {
	Entity string `form:"entity"`
	Query  string `form:"query"`
}

type SearchResult interface {
	ArtistSearchResult | LabelSearchResult
}

type SearchResults[T SearchResult] struct {
	Results []T `json:"results"`
}
