package queries

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
