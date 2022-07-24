
WITH cte AS (SELECT
    COUNT(CASE WHEN (duration/60) <= 5 THEN session_id END) AS total_1,
    COUNT(CASE WHEN (duration/60) <= 10 AND (duration/60) > 5 THEN session_id END) AS total_2,
    COUNT(CASE WHEN (duration/60) <= 15 AND (duration/60) > 10 THEN session_id END) AS total_3,
    COUNT(CASE WHEN (duration/60) >= 15 THEN session_id END) AS total_4,
    "[0-5>" AS bin_1,
    "[5-10>" AS bin_2,
    "[10-15>"AS bin_3,
    "15 or more"AS bin_4
FROM Sessions)



SELECT bin_1 AS bin, total_1 AS total FROM cte
UNION ALL
SELECT bin_2 AS bin, total_2 AS total FROM cte
UNION ALL
SELECT bin_3 AS bin, total_3 AS total FROM cte
UNION ALL
SELECT bin_4 AS bin, total_4 AS total FROM cte