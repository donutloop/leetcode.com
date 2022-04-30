SELECT
   country_name,
   CASE
      WHEN
         SUM(weather_state) / COUNT(*) <= 15
      THEN
         'Cold'
      WHEN
         SUM(weather_state) / COUNT(*) >= 25
      THEN
         'Hot'
      ELSE
         'Warm'
   END AS  weather_type
FROM
   Countries
   INNER JOIN
      Weather
      ON Weather.country_id = Countries.country_id
WHERE
   YEAR(day) = 2019
   AND
   MONTH(day) = 11
GROUP BY country_name;