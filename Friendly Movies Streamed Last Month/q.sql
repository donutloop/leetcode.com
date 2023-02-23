SELECT
  DISTINCT(title)
FROM
  TVProgram AS p
  JOIN Content AS c ON c.content_id = p.content_id
WHERE
  c.Kids_content = 'Y'
  AND c.content_type = 'Movies'
  AND MONTH(p.program_date) = 6
  AND YEAR(p.program_date) = 2020;