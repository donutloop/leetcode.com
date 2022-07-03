WITH cte (airport_id, flights_count_max) AS
         (
             SELECT
                 airport_id,
                 SUM(flights_count) AS flights_count_max
             FROM
                 (
                     SELECT
                         departure_airport AS airport_id,
                         flights_count
                     FROM
                         Flights
                     UNION ALL
                     SELECT
                         arrival_airport AS airport_id,
                         flights_count
                     FROM
                         Flights
                 )
                     AS Flights
             GROUP BY
                 airport_id
         )

SELECT
    airport_id
FROM
    cte
WHERE
        flights_count_max =
        (
            SELECT
                MAX(flights_count_max)
            FROM
                cte
        )
;
