SELECT
    ABS(b.x-c.x) AS shortest
FROM (SELECT
          a.x,
          ROW_NUMBER() OVER (
        ORDER BY
          a.x
      ) AS 'row_number'
      FROM
          Point AS a
     ) AS b
         INNER JOIN (
    SELECT
        a.x,
        ROW_NUMBER() OVER (
        ORDER BY
          a.x
      ) AS 'row_number'
    FROM
        Point AS a
)  AS c ON (c.row_number > b.row_number) ORDER BY  ABS(b.x-c.x) LIMIT 1



