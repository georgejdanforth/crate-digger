package queries

const SearchArtist = `SELECT
	artist.id               AS id,
	artist.gid              AS gid,
	artist.name             AS name,
	artist.sort_name        AS sort_name,
	artist.begin_date_year  AS begin_date_year,
	artist.begin_date_month AS begin_date_month,
	artist.begin_date_day   AS begin_date_day,
	artist.end_date_year    AS end_date_year,
	artist.end_date_month   AS end_date_month,
	artist.end_date_day     AS end_date_day,
	artist.type             AS type,
	artist.comment          AS comment,
	area.name               AS area_name,
	MAX(name_scores.score)  AS score
FROM (
	SELECT name, similarity(name, $1) AS score
	FROM (
		SELECT name              FROM artist       UNION ALL
		-- SELECT sort_name AS name FROM artist       UNION ALL
		SELECT name              FROM artist_alias
		-- SELECT sort_name AS name FROM artist_alias
	) names
	WHERE similarity(name, $1) > $2
	ORDER BY score DESC
	LIMIT 50
) name_scores
LEFT JOIN artist_alias AS alias ON (alias.name = name_scores.name)
LEFT JOIN area ON artist.area = area.id
INNER JOIN artist ON (alias.artist = artist.id OR name_scores.name = artist.name)
GROUP BY artist.id, artist.gid, artist.name, area.name
ORDER BY score DESC, artist.name
LIMIT 25;`

const SearchLabel = `SELECT
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
	area.name              AS area_name,
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
LEFT JOIN area ON label.area = area.id
GROUP BY label.id, label.gid, label.name, area.name
ORDER BY score DESC, label.name
LIMIT 50;`
