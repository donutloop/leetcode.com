package kata

SELECT * FROM cinema c WHERE c.description != "boring" AND c.id % 2 = 1 ORDER BY c.rating DESC;

