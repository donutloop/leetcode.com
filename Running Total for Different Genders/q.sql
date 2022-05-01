SELECT gender,
    DAY,
    SUM(score_points) OVER (PARTITION BY gender
    ORDER BY gender,
    DAY) AS total
FROM Scores
GROUP BY gender,
    DAY
ORDER BY gender ASC,
    DAY ASC