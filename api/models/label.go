package models

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
