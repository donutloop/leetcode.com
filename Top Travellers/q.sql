SELECT name, COALESCE(SUM(distance), 0) AS travelled_distance
FROM Users AS u
LEFT JOIN Rides AS r
ON r.user_id = u.id
GROUP BY u.id
ORDER BY travelled_distance DESC, name ASC;